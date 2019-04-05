import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import Link from '../components/Link';
import List from '../components/List';
import Carret from '../components/Carret';
import ActionCreators from '../state/actions';

class Backups extends Component {

  renderItem = (item) => (
    <Link to={'/node/' + item.name}>
      <Carret />
      {item.name}
    </Link>
  )

  isLoaded = () => this.props.container && this.props.container.contents;

  async componentDidMount() {
    if (!this.isLoaded()) {
      const { name } = this.props.match.params;
      this.props.actions.listBackupsInContainer(name);
    }
  }

  render() {
    if (!this.isLoaded()) {
      return <div></div>;
    }

    return (
      <div className="Backups">
        <h2>Contents of {this.props.container.name}</h2>
        <List
          data={this.props.container.contents}
          renderItem={this.renderItem}
        />
      </div>
    );
  }
}

const mapStateToProps = (state, props) => ({
  container: state.containers.find(c => c.name === props.match.params.name),
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Backups);
