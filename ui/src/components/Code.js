import React, { PureComponent } from 'react';

import './Code.less';

class Code extends PureComponent {

  render() {
    return (
      <pre className={`code ${this.props.dark ? 'dark' : ''}`}>
        {this.props.children}
      </pre>
    );
  }
}

export default Code;
