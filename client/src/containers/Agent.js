import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { Table, Tag, Modal } from 'antd';

import Loading from '../components/Loading';
import ActionCreators from '../state/actions';
import { formatDate } from '../helpers';

class Agent extends Component {

  state = {
    detailsIndex: 0,
    detailsVisible: false,
  }

  columns = [
    {
      title: 'Type',
      dataIndex: 'type',
      key: 'type',
      render: (type) => <Tag color={this.getColorByType(type)} key={type}>{type.toUpperCase()}</Tag>,
    },
    {
      title: 'UTC Date',
      dataIndex: 'timestamp',
      key: 'timestamp',
      render: formatDate,
    },
    {
      title: 'Actions',
      key: 'action',
      render: (_, item) => (
        <a onClick={() => this.setState({ detailsIndex: item.key, detailsVisible: true })}>
          Show details
        </a>
      ),
    }
  ]

  getColorByType = (type) => {
    switch (type) {
      case 'backup_start': return 'orange';
      case 'backup_done': return 'green';
      default: return 'volcano';
    }
  }

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
            <Table columns={this.columns} dataSource={this.props.agentMessages} />
            {this.props.agentMessages.length > 1 && (
              <Modal
                title='Action details'
                visible={this.state.detailsVisible}
                footer={null}
                onCancel={() => this.setState({ detailsVisible: false })}
                width='80%'
              >
                <pre>{JSON.stringify(this.props.agentMessages[this.state.detailsIndex], null, 2)}</pre>
              </Modal>
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
