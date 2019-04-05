import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux'
import { BrowserRouter as Router, Route } from 'react-router-dom';

import Containers from '../containers/Containers';
import Trees from '../containers/Trees';
import Backups from '../containers/Backups';
import ActionCreators from '../state/actions';

class App extends Component {

  componentDidMount() {
    this.props.actions.setupApp();
  }

  render() {
    return (
      <Router>
        <div className="App">
          <Route exact path='/' component={Containers} />
          <Route exact path='/node/:name' component={Trees} />
          <Route exact path='/node/:name/tree/:treeName' component={Backups} />
        </div>
      </Router>
    );
  }
}

const mapStateToProps = state => ({
  isSetup: state.app.isSetup,
});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
