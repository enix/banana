import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import List from '../components/List';
import Loading from '../components/Loading';
import ActionCreators from '../state/actions';

class Agent extends Component {

  renderMessage = (message) => (
    <pre>{JSON.stringify(message, null, 2)}</pre>
  )

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
      <div className="Agent">
        <h2>Agent {this.props.agent.cn} from {this.props.agent.organization}</h2>
        <h4>Configuration</h4>
        <pre>{JSON.stringify(this.props.agent.config, null, 2)}</pre>
        <h4>Actions history</h4>
        {!this.props.agentMessages ? <Loading /> : (
          <List
            data={this.props.agentMessages}
            renderItem={this.renderMessage}
          />
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

  if (state.agentsMessages && state.agentsMessages[org]) {
    loaded.agentMessages = state.agentsMessages[org][cn];
  }

  return loaded;
};

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Agent);
