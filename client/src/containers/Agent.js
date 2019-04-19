import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { Table, Tag, Modal, Divider, Icon, Button } from 'antd';

import JsonTable from '../components/JsonTable';
import Code from '../components/Code';
import Loading from '../components/Loading';
import ActionCreators from '../state/actions';
import {
  formatDate,
  formatSnakeCase,
  getTagColor,
  generateRestoreCmd,
} from '../helpers';

class Agent extends Component {

  state = {
    detailsIndex: 0,
    configVisible: false,
    logsVisible: false,
    restoreVisible: false,
  }

  columns = [
    {
      title: 'Type',
      dataIndex: 'type',
      key: 'type',
      render: (type) => (
        <Tag color={getTagColor(type)} key={type}>
          {formatSnakeCase(type)}
        </Tag>
      ),
    },
    {
      title: `Time (${localStorage.getItem('dateFormat')})`,
      dataIndex: 'timestamp',
      key: 'timestamp',
      render: formatDate,
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, item) => (
        <div>
          <a onClick={() => this.showConfig(item.key)}>Show config</a>
          <Divider type='vertical' />
          <a href={null} onClick={() => this.showCommand(item.key)}>Show command</a>
          
          {item.logs && (
            <span>
              <Divider type='vertical' />
              <a onClick={() => this.showLogs(item.key)}>Show logs</a>
            </span>
          )}
          
          {item.type === 'backup_done' && (
            <span>
              <Divider type='vertical' />
              <a
                href={`https://console.nxs.enix.io/project/containers/container/${item.config.bucket}/${item.command.name}`}
                target='_blank'
                rel='noopener noreferrer'
              >
                <Icon type='link' /> View on storage
              </a>
              <Button style={{ float: 'right' }} onClick={() => this.showRestore(item.key)}>
                Restore
              </Button>
            </span>
          )}
        </div>
      ),
    }
  ]

  showConfig = detailsIndex => this.setState({ detailsIndex, configVisible: true })

  showCommand = detailsIndex => this.setState({ detailsIndex, commandVisible: true })

  showLogs = detailsIndex => this.setState({ detailsIndex, logsVisible: true })

  showRestore = detailsIndex => this.setState({ detailsIndex, restoreVisible: true })

  componentDidMount() {
    const { org, cn } = this.props.match.params;
    this.props.actions.getAgent(org, cn);
    this.props.actions.getAgentMessages(org, cn);
  }

  render() {
    if (!this.props.agent) {
      return <Loading />;
    }

    return (
      <div className='Agent'>
        <h2>Actions history for {this.props.agent.cn} from {this.props.agent.organization}</h2>
        {!this.props.agentMessages ? <Loading /> : (
          <div>
            <Table
              columns={this.columns}
              dataSource={this.props.agentMessages}
            />
            {this.props.agentMessages.length > 1 && (
              <div>
                <Modal
                  title='Action config'
                  visible={this.state.configVisible}
                  footer={null}
                  onCancel={() => this.setState({ configVisible: false })}
                  width='80%'
                >
                  <JsonTable data={this.props.agentMessages[this.state.detailsIndex].config} />
                </Modal>
                <Modal
                  title='Action command'
                  visible={this.state.commandVisible}
                  footer={null}
                  onCancel={() => this.setState({ commandVisible: false })}
                >
                  <JsonTable data={this.props.agentMessages[this.state.detailsIndex].command} />
                </Modal>
                <Modal
                  title='Action logs'
                  visible={this.state.logsVisible}
                  footer={null}
                  onCancel={() => this.setState({ logsVisible: false })}
                  width='80%'
                >
                  <Code dark>
                    {this.props.agentMessages[this.state.detailsIndex].logs}
                  </Code>
                </Modal>
                <Modal
                  title='Restore backup'
                  visible={this.state.restoreVisible}
                  footer={null}
                  onCancel={() => this.setState({ restoreVisible: false })}
                >
                  <Code dark>
                    {generateRestoreCmd(this.props.agentMessages[this.state.detailsIndex])}                    
                  </Code>
                </Modal>
              </div>
            )}
          </div>
        )}
      </div>
    );
  }
}

const mapStateToProps = (state, props) => {
  const loaded = {};
  const { org, cn } = props.match.params;

  if (state.agents && state.agents[org]) {
    loaded.agent = state.agents[org][cn];
  }

  if (state.agentsMessages && state.agentsMessages[org] && state.agentsMessages[org][cn]) {
    loaded.agentMessages = state.agentsMessages[org][cn].map((msg, key) => ({ ...msg, key }));
  }

  return loaded;
};

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Agent);
