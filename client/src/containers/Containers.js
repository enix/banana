import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'

import Link from '../components/Link';
import List from '../components/List';
import ActionCreators from '../state/actions';

class Containers extends Component {

  constructor({ match }) {
    super();
    console.log(match);
  }

  componentDidMount() {
    this.props.actions.listBackupContainers();
  }

  render() {
    return (
      <div className="Containers">
        <List
          data={this.props.containers}
          renderItem={item => <Link to={'/containers/' + item.name}>{item.name}</Link>}
        />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  containers: state.containers,
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Containers);
