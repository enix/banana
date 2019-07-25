import React, { Component } from 'react';
import { connect } from 'react-redux';
import { mapDispatchToProps } from 'redux-saga-wrapper';
import {
  Table,
  Tag,
  Divider,
  Icon,
  Button,
  Switch,
  Row,
  Col,
} from 'antd';

import MessageModals from './MessageModals';
import Loading from '../components/Loading';
import {
  formatDate,
  getSelectedTimezoneName,
  formatSnakeCase,
  getTagColor,
} from '../helpers';

class MessagesList extends Component {

  state = {
    detailsIndex: 0,
    configVisible: false,
    commandVisible: false,
    logsVisible: false,
    restoreVisible: false,
    actionsStartVisible: false,
    routinesVisible: false,
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
      title: `Time (${getSelectedTimezoneName()})`,
      dataIndex: 'timestamp',
      key: 'timestamp',
      render: formatDate,
    },
    // {
    //   title: `Backup name`,
    //   dataIndex: 'command.name',
    //   key: 'name',
    //   render: name => name ? name : '-',
    // },
    // {
    //   title: `Backup type`,
    //   dataIndex: 'command.type',
    //   key: 'backup_type',
    //   render: (type) => !type ? '-' : (
    //     <Tag color={getTypeTagColor(type)} key={type}>
    //       {formatSnakeCase(type)}
    //     </Tag>
    //   ),
    // },
    {
      title: 'Actions',
      key: 'action',
      render: (_, item) => (
        <div>
          <Button type='link' onClick={() => this.showConfig(item.key)}>Show config</Button>
          <Divider type='vertical' />
          <Button type='link' onClick={() => this.showCommand(item.key)}>Show command</Button>

          {item.logs && (
            <span>
              <Divider type='vertical' />
              <Button type='link' onClick={() => this.showLogs(item.key)}>Show logs</Button>
            </span>
          )}

          {item.type === 'backup_done' && (
            <span>
              <Divider type='vertical' />
              <a
                href={`https://banana.dev.enix.io/api/agents/${this.props.agentID}/messages/${item.timestamp}/artifacts.gzip`}
                target='_blank'
                rel='noopener noreferrer'
              >
                <Icon type='link' /> Download artifacts
              </a>
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

  toggleActionsStart = actionsStartVisible => this.setState({ actionsStartVisible });

  toggleRoutines = routinesVisible => this.setState({ routinesVisible });

  getAgentMessages = () => {
    let actions = this.props.agentMessages;

    if (!this.state.actionsStartVisible) {
      actions = actions.filter(message => !/.*start.*/gi.test(message.type));
    }

    if (!this.state.routinesVisible) {
      actions = actions.filter(message => !/.*routine.*/gi.test(message.type));
    }

    return actions;
  };

  render() {
    return (
      <div className='MessagesList'>
        <Row>
          <Col style={{ textAlign: 'right' }}>
            Display actions start
            <Switch onChange={this.toggleActionsStart} style={{ marginLeft: 10 }} />
            <Divider type='vertical' />
            Display routines
            <Switch onChange={this.toggleRoutines} style={{ marginLeft: 10 }} />
          </Col>
        </Row>
        {!this.props.agentMessages ? <Loading /> : (
          <div>
            <Table
              columns={this.columns}
              dataSource={this.getAgentMessages()}
            />
            {this.props.agentMessages.length > 0 && (
              <MessageModals
                agentMessages={this.props.agentMessages}
                detailsIndex={this.state.detailsIndex}
                configVisible={this.state.configVisible}
                logsVisible={this.state.logsVisible}
                restoreVisible={this.state.restoreVisible}
                commandVisible={this.state.commandVisible}
                onCancel={() => this.setState({
                  restoreVisible: false,
                  logsVisible: false,
                  commandVisible: false,
                  configVisible: false,
                })}
              />
            )}
          </div>
        )}
      </div>
    );
  }
}

const mapStateToProps = (state, props) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(MessagesList);
