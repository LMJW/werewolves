import React from 'react';
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

interface LoginProps {}

const Login: React.FC<LoginProps> = props => {
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

export default Login;
