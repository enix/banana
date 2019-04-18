import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import List from '../components/List';
import Loading from '../components/Loading';
import ActionCreators from '../state/actions';
import { formatDate } from '../helpers';

class Agent extends Component {

  renderMessage = (message, key) => (
    <div>
      <h4 style={{ display: 'inline-block', marginRight: 20 }}>
        {formatDate(message.timestamp)} -
        <b> {message.type}</b>
      </h4>
      <a
        href={`#collapseExample-${key}`}
        data-toggle='collapse'
        aria-expanded='false'
        aria-controls={`collapseExample-${key}`}
      >
        Toggle details
      </a>
      <div className='collapse' id={`collapseExample-${key}`}>
        <div class='card card-body'>
          <pre style={{ overflow: 'hidden', textOverflow: 'ellipsis' }}>
            {JSON.stringify(message, null, 2)}
          </pre>
        </div>
      </div>
    </div>
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
      <div className='Agent'>
        <h2>Actions history for {this.props.agent.cn} from {this.props.agent.organization}</h2>
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
