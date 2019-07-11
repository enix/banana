import React, { Component } from 'react';
import { connect } from 'react-redux';
import { mapDispatchToProps } from 'redux-saga-wrapper';
import { Col } from 'antd';

import MessagesList from './MessagesList';
import BackupsList from './BackupsList';
import Loading from '../components/Loading';

class Agents extends Component {

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
        <h1 style={{ marginBottom: 30 }}>
          {this.props.agent.organization} /
          <strong> {this.props.agent.cn}</strong>
        </h1>

        <div>
          <Col span={18}>
            <h3>Backup list</h3>
          </Col>
          <BackupsList agentMessages={this.props.agentMessages} />
        </div>

        <div style={{ marginTop: 30 }}>
          <Col span={18} >
            <h3>Messages history</h3>
          </Col>
          <MessagesList agentMessages={this.props.agentMessages} />
        </div>
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

export default connect(mapStateToProps, mapDispatchToProps)(Agents);
