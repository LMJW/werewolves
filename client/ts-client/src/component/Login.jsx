import React from 'react';
import {
	FormControl,
	Paper,
	InputLabel,
	Input,
	InputAdornment,
	Button,
	Typography,
	CircularProgress
} from '@material-ui/core';
import useStyles from '../style/styles';
import { connect } from 'react-redux';
import actionTypes from '../store/constants';

const Login = props => {
	const { onLoginFormChange, onLoginSubmit, username, loading, error } = props;
	const classes = useStyles();
	return (
		<div className={classes.container}>
			<Paper className={classes.paper}>
				{loading && <CircularProgress />}
				{!loading &&
					error === null && (
						<div>
							<h2>{'Login'}</h2>
							<Typography>
								{
									'Notice: this is a temporal account, so you do not need to enter a password. Your account info will be deleted after few hours.'
								}
							</Typography>
							<FormControl required={true} fullWidth={true} className={classes.field}>
								<InputLabel htmlFor="user">Name you want to be displayed in game</InputLabel>
								<Input
									id="username"
									startAdornment={<InputAdornment position="start" />}
									value={username}
									onChange={e => {
										onLoginFormChange(e.target.value);
									}}
								/>
							</FormControl>
							<Button
								type="submit"
								fullWidth
								variant="contained"
								color="primary"
								className={classes.submit}
								onClick={e => {
									onLoginSubmit(username);
								}}
							>
								Sign In
							</Button>
						</div>
					)}
				{!loading && error !== null && <div>{error}</div>}
			</Paper>
		</div>
	);
};

const mapStateToProps = state => ({
	username: state.username
});

const mapDispatchToProps = dispatch => ({
	onLoginSubmit: username => dispatch({ type: actionTypes.USER_LOGIN_DONE, username }),
	onLoginFormChange: username => dispatch({ type: actionTypes.UPDATE_LOGIN_FORM_CONTEXT, username })
});

export default connect(mapStateToProps, mapDispatchToProps)(Login);
