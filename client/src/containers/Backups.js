import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import Link from '../components/Link';
import List from '../components/List';
import Carret from '../components/Carret';
import ActionCreators from '../state/actions';

class Backups extends Component {

  renderItem = (item) => (
    <Link to={window.location.pathname + '/restore/' + item.time}>
      <Carret />
      {item.time}
    </Link>
  )

  isLoaded = () => this.props.tree && this.props.tree.contents;

  async componentDidMount() {
    if (!this.isLoaded()) {
      const { name, treeName } = this.props.match.params;
      this.props.actions.listBackupsForTree(name, treeName);
    }
  }

  render() {
    if (!this.isLoaded()) {
      return <div></div>;
    }

    return (
      <div className="Backups">
        <h2>Contents of {this.props.tree.name}</h2>
        <List
          data={this.props.tree.contents}
          renderItem={this.renderItem}
        />
      </div>
    );
  }
}

const mapStateToProps = (state, props) => {
  const container = state.containers.find(c => c.name === props.match.params.name);
  if (!container) {
    return {};
  }

  return {
    tree: container.contents.find(t => t.name === props.match.params.treeName),
  };
};

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Backups);
