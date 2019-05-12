import React from 'react';
import { RouteComponentProps } from 'react-router-dom';
import {
  FormControl,
  Paper,
  InputLabel,
  Input,
  InputAdornment,
  Button,
  Typography
} from '@material-ui/core';
import useStyles from '../style/styles';
import { connect } from 'react-redux';
import { UserLoginState } from '../store/login/types';

interface LoginProps {
  loading: boolean;
  error?: string;
}

type Props = LoginProps & UserLoginState;

const Login = (props: LoginProps) => {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <Paper className={classes.paper}>
        <h2>{'Login'}</h2>
        <Typography>
          {
            'Notice: this is a temporal account, so you do not need to enter a password. Your account info will be deleted after few hours.'
          }
        </Typography>
        <FormControl required={true} fullWidth={true} className={classes.field}>
          <InputLabel htmlFor='user'>
            Name you want to be displayed in game
          </InputLabel>
          <Input
            id='username'
            startAdornment={<InputAdornment position='start' />}
          />
        </FormControl>
        <Button
          type='submit'
          fullWidth
          variant='contained'
          color='primary'
          className={classes.submit}
        >
          Sign In
        </Button>
      </Paper>
    </div>
  );
};

const mapStateToProps = (state: UserLoginState) => ({
  loginState: state.loginState,
  username: state.username,
  session: state.session
});

export default connect(mapStateToProps)(Login);
