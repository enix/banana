import React, { Component } from 'react';

import './Code.less';

class Code extends Component {
  
  render() {
    return (
      <pre className={`code ${this.props.dark ? 'dark' : ''}`}>
        {this.props.children}
      </pre>
    );
  }
}

export default Code;
