import React, { Component } from 'react';
import HomePage from '../components/home';
import Header from '../components/header';
import { orange700, orange300, orange900, brown700, brown400, brown900, grey900, grey50 } from 'material-ui/styles/colors'
import FlatButton from 'material-ui/FlatButton'
import getMuiTheme from 'material-ui/styles/getMuiTheme'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import injectTapEventPlugin from 'react-tap-event-plugin';

try {
	injectTapEventPlugin();
} catch(e) {}

const styles = {
	container: {
		textAlign: 'center',
		paddingTop: 200
	}
}

const muiTheme = {
	palette: {
		primary1Color: orange700,
		primary2Color: orange300,
		primary3Color: orange900,
		accent1Color: brown700,
		accent2Color: brown400,
		accent3Color: brown900,
		textColor: grey50,
		alternateTextColor: grey50,
	}
}

export default class extends Component {

	static async getInitialProps ({ req }) {
		const userAgent = req ? req.headers['user-agent'] : navigator.userAgent;
		return { userAgent };
	}

	render() {
		return (
			<MuiThemeProvider muiTheme={getMuiTheme({userAgent: this.props.userAgent, ...muiTheme})}>
			<div>
				<Header />
				<h1 className={'logo'}>Welcome to direct-store!</h1>
				<HomePage />
				<FlatButton
					label='Ok'
					primary={Boolean(true)}
					onTouchTap={this.handleRequestClose}
				/>

				<style jsx global>{` 
					body {
						font-family: sans-serif;
						margin: 0;
						padding: 0;
					}
				`}</style>
				<style jsx>{`
					h1 {
						color: blue;
					}
				`}</style>
			</div>
			</MuiThemeProvider>
		);
	}
};
