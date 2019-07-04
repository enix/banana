import React, { Component } from 'react';
import { connect } from 'react-redux';
import { mapDispatchToProps } from 'redux-saga-wrapper';
import {
  Form,
  Input,
	Icon,
	Button,
	Tooltip,
	InputNumber,
} from 'antd';

class DuplicityConfiguration extends Component {

	nextFieldUniqueId = 0

	getFormKey = key => `duplicity[${this.props._key}].${key}s`

	add = subkey => {
		const { form } = this.props;
		const formKey = `duplicity[${this.props._key}].${subkey}-keys`;
    const keys = form.getFieldValue(formKey).concat(++this.nextFieldUniqueId);
    form.setFieldsValue({ [formKey]: keys });
  }

  remove = (subkey, key) => {
    const { form } = this.props;
		const formKey = `duplicity[${this.props._key}].${subkey}-keys`;
		console.log(formKey)
    form.setFieldsValue({
      [formKey]: form.getFieldValue(formKey).filter(k => k !== key),
    });
	}

	renderField = (type, key, index) => {
		const { getFieldDecorator } = this.props.form;

		return <Form.Item label={`${type[0].toUpperCase()}${type.slice(1)} directory`} key={index}>
			{getFieldDecorator(`${this.getFormKey(type)}[${key}]`, {
				validateTrigger: ['onChange', 'onBlur'],
				rules: [
					{
						required: true,
						whitespace: true,
						message: `Please enter an ${type} rule`,
					},
				],
			})(<Input placeholder={type === 'include' ? '/' : '/proc'} style={{ width: this.props.formWidthSm, marginRight: 8 }} />)}
			<Icon type='minus-circle-o' onClick={() => this.remove(type, key)} />
		</Form.Item>;
	}

  render() {
		const { getFieldValue, getFieldDecorator } = this.props.form;
		const { _key: key } = this.props;

		getFieldDecorator(`duplicity[${key}].include-keys`, { initialValue: [0], required: false });
		getFieldDecorator(`duplicity[${key}].exclude-keys`, { initialValue: [], required: false });
		const includeItems = getFieldValue(`duplicity[${key}].include-keys`).map((key, index) => this.renderField('include', key, index));
		const excludeItems = getFieldValue(`duplicity[${key}].exclude-keys`).map((key, index) => this.renderField('exclude', key, index));

    return <>
			<Form.Item label={
				<Tooltip title='Full backups interval, in number of backups. Setting this value to 2 will result in half of the backups being fulls and the other half being incrementals.'>
					Full backup every <Icon type='question-circle' />
				</Tooltip>
			}>
				{getFieldDecorator(`duplicity[${key}].fullEvery`, {
					validateTrigger: ['onChange', 'onBlur'],
					rules: [{
						required: true,
						message: 'Please choose the full backups interval.',
					}],
				})(<InputNumber placeholder='7' style={{ width: this.formWidthSm, marginRight: 8 }} />)}
			</Form.Item>

			{includeItems}
			{excludeItems}

			<Form.Item style={{ display: 'inline-block', marginLeft: '25%', marginRight: 20 }}>
				<Button type='dashed' onClick={() => this.add('include')}>
					<Icon type='plus' /> Add include rule
				</Button>
			</Form.Item>
			<Form.Item style={{ display: 'inline-block' }}>
				<Button type='dashed' onClick={() => this.add('exclude')}>
					<Icon type='plus' /> Add exclude rule
				</Button>
			</Form.Item>
		</>;
  }
}

const mapStateToProps = () => ({});

export default connect(mapStateToProps, mapDispatchToProps)(DuplicityConfiguration);
