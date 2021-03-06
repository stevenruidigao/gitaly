require 'spec_helper'

describe Gitlab::Git::SshAuth do
  describe '.from_gitaly' do
    it 'initializes based on ssh_key and known_hosts in the request' do
      result = described_class.from_gitaly(double(ssh_key: 'SSH KEY', known_hosts: 'KNOWN HOSTS'))

      expect(result.class).to eq(described_class)
      expect(result.ssh_key).to eq('SSH KEY')
      expect(result.known_hosts).to eq('KNOWN HOSTS')
    end
  end

  describe '#setup' do
    subject { described_class.new(ssh_key, known_hosts).setup { |env| env } }

    context 'no credentials' do
      let(:ssh_key) { nil }
      let(:known_hosts) { nil }

      it 'writes no tempfiles' do
        expect(Tempfile).not_to receive(:new)

        is_expected.to eq({})
      end
    end

    context 'just the SSH key' do
      let(:ssh_key) { 'Fake SSH key' }
      let(:known_hosts) { nil }

      it 'writes the SSH key file' do
        ssh_key_file = stub_tempfile('/tmp files/keyFile', 'gitlab-shell-key-file', chmod: 0o400)

        is_expected.to eq(build_env(ssh_key_file: ssh_key_file.path))

        expect(ssh_key_file.string).to eq(ssh_key)
      end
    end

    context 'just the known_hosts file' do
      let(:ssh_key) { nil }
      let(:known_hosts) { 'Fake known_hosts data' }

      it 'writes the known_hosts file and script' do
        known_hosts_file = stub_tempfile('/tmp files/knownHosts', 'gitlab-shell-known-hosts', chmod: 0o400)

        is_expected.to eq(build_env(known_hosts_file: known_hosts_file.path))

        expect(known_hosts_file.string).to eq(known_hosts)
      end
    end

    context 'SSH key and known_hosts file' do
      let(:ssh_key) { 'Fake SSH key' }
      let(:known_hosts) { 'Fake known_hosts data' }

      it 'writes SSH key, known_hosts and script files' do
        ssh_key_file = stub_tempfile('id_rsa', 'gitlab-shell-key-file', chmod: 0o400)
        known_hosts_file = stub_tempfile('known_hosts', 'gitlab-shell-known-hosts', chmod: 0o400)

        is_expected.to eq(build_env(ssh_key_file: ssh_key_file.path, known_hosts_file: known_hosts_file.path))

        expect(ssh_key_file.string).to eq(ssh_key)
        expect(known_hosts_file.string).to eq(known_hosts)
      end
    end
  end

  def build_env(ssh_key_file: nil, known_hosts_file: nil)
    opts = []

    if ssh_key_file
      opts << %('-oIdentityFile="#{ssh_key_file}"')
      opts << %q('-oIdentitiesOnly="yes"')
    end

    if known_hosts_file
      opts << %q('-oStrictHostKeyChecking="yes"')
      opts << %('-oUserKnownHostsFile="#{known_hosts_file}"')
    end

    { 'GIT_SSH_COMMAND' => %(ssh #{opts.join(' ')}) }
  end

  def stub_tempfile(name, filename, chmod:)
    file = StringIO.new

    allow(file).to receive(:path).and_return(name)

    expect(Tempfile).to receive(:new).with(filename).and_return(file)
    expect(file).to receive(:chmod).with(chmod)
    expect(file).to receive(:close!)

    file
  end
end
