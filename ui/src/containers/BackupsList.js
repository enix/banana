import React, { Component } from 'react';
import { connect } from 'react-redux';
import { mapDispatchToProps } from 'redux-saga-wrapper';
import filesize from 'filesize';
import {
  Table,
  Tag,
  Divider,
  Icon,
  Button,
} from 'antd';

import MessageModals from './MessageModals';
import Loading from '../components/Loading';
import {
  formatDate,
  getSelectedTimezoneName,
  formatSnakeCase,
  getTypeTagColor,
} from '../helpers';

class BackupsList extends Component {

  state = {
    detailsIndex: 0,
    configVisible: false,
    logsVisible: false,
    restoreVisible: false,
  }

  columns = [
    {
      title: 'Name',
      dataIndex: 'command.name',
      key: 'name',
      render: name => name ? name : '-',
    },
    {
      title: 'Plugin',
      dataIndex: 'config.plugin',
      key: 'plugin',
      render: name => name ? name : '-',
    },
    {
      title: `Time (${getSelectedTimezoneName()})`,
      dataIndex: 'timestamp',
      key: 'timestamp',
      render: formatDate,
    },
    {
      title: 'Type (last)',
      dataIndex: 'command.type',
      key: 'backup_type',
      render: (type) => !type ? '-' : (
        <Tag color={getTypeTagColor(type)} key={type}>
          {formatSnakeCase(type)}
        </Tag>
      ),
    },
    {
      title: 'Size (last)',
      dataIndex: 'metadata.size',
      key: 'size',
      render: size => size ? filesize(size) : '-',
    },
    {
      title: 'Size (total)',
      dataIndex: 'metadata.totalSize',
      key: 'total_size',
      render: size => size ? filesize(size) : '-',
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, item) => (
        <div>
          <Button type='link' onClick={() => this.showConfig(item.key)}>Show config</Button>
          {/* <Divider type='vertical' />
          <Button type='link' onClick={() => this.showCommand(item.key)}>Show command</Button> */}

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

  getBackupList = () => this.props.agentMessages.reduce((acc, it) => {
    if (it.type !== 'backup_done') {
      return acc;
    }

    const mostRecent = acc.find(b => b.command.name === it.command.name);
    if (!mostRecent) {
      it.metadata.totalSize = it.metadata.size;
      acc.push(it);
    }
    else {
      mostRecent.metadata.totalSize += it.metadata.size || 0;
    }

    return acc;
  }, []);

  render() {
    return (
      <div className='BackupsList'>
        {!this.props.agentMessages ? <Loading /> : (
          <div>
            <Table
              columns={this.columns}
              dataSource={this.getBackupList()}
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

export default connect(mapStateToProps, mapDispatchToProps)(BackupsList);
