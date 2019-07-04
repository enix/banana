import React, { Component } from 'react';
import { connect } from 'react-redux';
import { mapDispatchToProps } from 'redux-saga-wrapper';
import CopyToClipboard from 'react-copy-to-clipboard';
import {
  Form,
  Input,
  InputNumber,
  Divider,
  Icon,
  Button,
  Select,
  Tooltip,
  Modal,
} from 'antd';

import Code from '../components/Code';
import DuplicityConfig from './plugins/DuplicityConfig';

class Configuration extends Component {

  state = {
    result: '',
    resultVisible: false,
  }

  plugins = {
    duplicity: {
      name: 'Duplicity',
      component: DuplicityConfig,
    },
  }

  formItemLayout = {
    labelCol: {
      xs: { span: 24 },
      sm: { span: 6 },
    },
    wrapperCol: {
      xs: { span: 24 },
      sm: { span: 8 },
    },
  }

  formItemLayoutWithOutLabel = {
    wrapperCol: {
      xs: { span: 24, offset: 0 },
      sm: { span: 8, offset: 6 },
    },
  }

  nextFieldUniqueId = 0

  formWidthSm = 420

  componentDidMount() {
    this.add();
  }

  generateSchedule = values => values.keys.reduce((result, key) => {
    const plugin = values.plugin[key];

    result[values.name[key]] = {
      interval: values.interval[key],
      plugin,
      ...this.plugins[plugin].component.generateSchedule(values[plugin][key]),
    };

    return result;
  }, {});

  handleSubmit = evt => {
    evt.preventDefault();

    this.props.form.validateFields((err, values) => {
      // if (!err) {
        this.setState({
          result: JSON.stringify(this.generateSchedule(values), null, 2),
          resultVisible: true,
        });
      // }
    });
  }

  add = () => {
    const { form } = this.props;
    const keys = form.getFieldValue('keys').concat(++this.nextFieldUniqueId);
    form.setFieldsValue({ keys });
  }

  remove = key => {
    const { form } = this.props;
    form.setFieldsValue({
      keys: form.getFieldValue('keys').filter(k => k !== key),
    });
  }

  render() {
    const { getFieldDecorator, getFieldValue } = this.props.form;
    getFieldDecorator('keys', { initialValue: [] });

    const items = getFieldValue('keys').map((key) => {
      const pluginConfiguration = this.plugins[getFieldValue(`plugin[${key}]`)] || null;

      return (
        <div key={key}>
          <Form.Item label='Backup name'>
            {getFieldDecorator(`name[${key}]`, {
              validateTrigger: ['onChange', 'onBlur'],
              rules: [{
                required: true,
                whitespace: true,
                message: 'Please choose a backup name.',
              }],
            })(<Input placeholder='Root filesystem' style={{ width: this.formWidthSm, marginRight: 8 }} />)}
            <Icon type='minus-circle-o' onClick={() => this.remove(key)} />
          </Form.Item>

          <Form.Item label={
            <Tooltip title='Interval between backups, in days'>
              Interval <Icon type='question-circle' />
            </Tooltip>
          }>
            {getFieldDecorator(`interval[${key}]`, {
              validateTrigger: ['onChange', 'onBlur'],
              rules: [{
                required: true,
                message: 'Please choose an interval.',
              }],
            })(<InputNumber placeholder='0.042' style={{ width: this.formWidthSm, marginRight: 8 }} step={0.1} />)}
          </Form.Item>

          <Form.Item label='Plugin'>
            {getFieldDecorator(`plugin[${key}]`, {
              validateTrigger: ['onChange', 'onBlur'],
              rules: [{
                required: true,
                message: 'Please choose a plugin.',
              }],
            })(
              <Select
                showSearch
                style={{ width: this.formWidthSm }}
                placeholder='Select a plugin'
                optionFilterProp='children'
                filterOption={(input, option) => option.props.children.toLowerCase().indexOf(input.toLowerCase()) !== -1}
              >
                {Object.keys(this.plugins).map(plugin => (
                  <Select.Option value={plugin} key={plugin}>{this.plugins[plugin].name}</Select.Option>
                ))}
              </Select>
            )}
          </Form.Item>


          {pluginConfiguration && <>
            <div style={{ marginLeft: '25%', width: this.formWidthSm }}>
              <Divider>Configure {pluginConfiguration.name}</Divider>
            </div>
            <pluginConfiguration.component
              form={this.props.form}
              formWidthSm={this.formWidthSm}
              formItemLayoutWithOutLabel={this.formItemLayoutWithOutLabel}
              _key={key}
            />
          </>}

          <Divider />
        </div>
      );
    });

    return (
      <>
        <h1>Create a schedule for your agents</h1>
        <Divider />
        <Form {...this.formItemLayout} onSubmit={this.handleSubmit}>
          {items}
          <Form.Item style={{ display: 'inline-block', marginLeft: '25%', marginRight: 20 }}>
            <Button type='dashed' onClick={this.add}>
              <Icon type='plus' /> Add backup
            </Button>
          </Form.Item>
          <Form.Item style={{ display: 'inline-block' }}>
            <Button type='primary' htmlType='submit'>
              Generate schedule
            </Button>
          </Form.Item>
        </Form>

        <Modal
          title='Generated configuration'
          width='40%'
          visible={this.state.resultVisible}
          onCancel={() => this.setState({ resultVisible: false })}
          footer={[
            <Button key='back' onClick={() => this.setState({ resultVisible: false })}>
              Close
            </Button>,
            <CopyToClipboard key='copy' text={this.state.result}>
              <Button type='primary'>
                Copy
              </Button>
            </CopyToClipboard>,
            <Button
              key='download'
              type='primary'
              href={`data:application/octet-stream,${encodeURIComponent(this.state.result)}`}
              style={{ marginLeft: 5 }}
            >
              Download
            </Button>,
          ]}
        >
          <Code dark id='result'>{this.state.result}</Code>
        </Modal>
      </>
    );
  }
}

const mapStateToProps = () => ({});

export default connect(mapStateToProps, mapDispatchToProps)(Form.create({ name: 'configuration' })(Configuration));
