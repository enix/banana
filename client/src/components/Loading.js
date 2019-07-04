import React, { PureComponent } from 'react';
import { Icon } from 'antd';

class Loading extends PureComponent {

  render() {
    return (
      <div style={{ width: '100%', textAlign: 'center' }}>
        <Icon spin type="loading-3-quarters" />
      </div>
    );
  }
}

export default Loading;
