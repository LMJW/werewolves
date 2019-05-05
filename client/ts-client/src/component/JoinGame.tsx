import * as React from 'react';
import useStyles from '../style/styles';
import {
  Paper,
  Button,
  FormControl,
  Input,
  InputLabel,
  InputAdornment
} from '@material-ui/core';

interface JoinProps {}

const JoinGame: React.FC<JoinProps> = props => {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <Paper className={classes.paper}>
        <h2>Room ID</h2>
        <FormControl required={true} fullWidth={true} className={classes.field}>
          <InputLabel htmlFor='room id'>
            Enter the room id you want to join
          </InputLabel>
          <Input
            id='roomid'
            startAdornment={<InputAdornment position='start' />}
          />
        </FormControl>
        <div className={classes.field}>
          <Button
            type='submit'
            variant='contained'
            color='primary'
            className={classes.submit}
          >
            Join
          </Button>
          <Button
            type='submit'
            variant='contained'
            color='primary'
            className={classes.submit}
          >
            Back
          </Button>
        </div>
      </Paper>
    </div>
  );
};

export default JoinGame;
