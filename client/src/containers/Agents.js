import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { Table, Divider, Tag } from 'antd';

import Loading from '../components/Loading';
import ActionCreators from '../state/actions';

class Agents extends Component {

  columns = [
    {
      title: 'Name',
      dataIndex: 'cn',
      key: 'cn',
      render: (text, item) => <Link to={`/agent/${item.organization}/${item.cn}`}>{text}</Link>,
    },
    {
      title: 'Organization',
      dataIndex: 'organization',
      key: 'organization',
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status',
      render: tags => (
        <span>
          {tags.map((tag, key) => <Tag color={tag.color} key={key}>{tag.message.toUpperCase()}</Tag>)}
        </span>
      ),
    },
  ]

  componentDidMount() {
    this.props.actions.listAgents();
  }

  render() {
    if (!this.props.agents) {
      return <Loading center />
    }

    return (
      <div className='Agents'>
        <h1>Agents</h1>
        <Divider />
        <Table columns={this.columns} dataSource={this.props.agents} />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  agents: state.agentList && state.agentList.map((agent, key) => ({ ...agent, key, status: [
    {
      color: 'green',
      message: 'up',
    },
  ]})),
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Agents);
