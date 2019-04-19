import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { Form, Divider, Select } from 'antd';

import ActionCreators from '../state/actions';

class Settings extends Component {

  state = {
    fields: {
      dateFormat: localStorage.getItem('dateFormat') || 'local'
    }
  }

  componentDidMount() {
    this.props.form.validateFields();
  }

  onValueChange = (key, value) => localStorage.setItem(key, value)

  render() {
    return (
      <div>
        <h2>Local settings</h2>
        <Divider />
        <span>Preferred date format</span>
        <Select
          onChange={evt => this.onValueChange('dateFormat', evt)}
          defaultValue={this.state.fields.dateFormat}
          style={{ width: 100, marginLeft: 50 }}
        >
          <Select.Option value='UTC'>UTC</Select.Option>
          <Select.Option value='local'>Local</Select.Option>
        </Select>
      </div>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(ActionCreators, dispatch),
});

export default connect(mapStateToProps, mapDispatchToProps)(Form.create({ name: 'settings' })(Settings));
