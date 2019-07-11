import React, { Component } from 'react';
import { connect } from 'react-redux';
import { mapDispatchToProps } from 'redux-saga-wrapper';
import { Modal } from 'antd';

import JsonTable from '../components/JsonTable';
import Code from '../components/Code';
import { generateRestoreCmd } from '../helpers';

class MessageModals extends Component {

  render() {
    return (
			<div>
				<Modal
					title='Action config'
					visible={this.props.configVisible}
					footer={null}
					onCancel={this.props.onCancel}
					width='80%'
				>
					<JsonTable data={this.props.agentMessages[this.props.detailsIndex].config} />
				</Modal>
				<Modal
					title='Action command'
					visible={this.props.commandVisible}
					footer={null}
					onCancel={this.props.onCancel}
				>
					<JsonTable data={this.props.agentMessages[this.props.detailsIndex].command} />
				</Modal>
				<Modal
					title='Action logs'
					visible={this.props.logsVisible}
					footer={null}
					onCancel={this.props.onCancel}
					width='80%'
				>
					<Code dark>
						{this.props.agentMessages[this.props.detailsIndex].logs}
					</Code>
				</Modal>
				<Modal
					title='Restore backup'
					width='40%'
					visible={this.props.restoreVisible}
					footer={null}
					onCancel={this.props.onCancel}
				>
					<Code dark>
						{generateRestoreCmd(this.props.agentMessages[this.props.detailsIndex])}
					</Code>
				</Modal>
			</div>
    );
  }
}

const mapStateToProps = (state, props) => ({});

export default connect(mapStateToProps, mapDispatchToProps)(MessageModals);
