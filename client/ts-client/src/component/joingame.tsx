import * as React from 'react';
import useStyles from '../style/styles';
import { Paper, Typography, Button } from '@material-ui/core';

interface UserProps {}

const JoinGame: React.FC<UserProps> = props => {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <Paper className={classes.paper}>
        <Typography className={classes.actions}>Do you want to </Typography>
        <Button
          className={classes.submit}
          color='primary'
          type='submit'
          variant='contained'
        >
          Start a new game
        </Button>
        <Typography className={classes.actions}>or</Typography>
        <Button
          className={classes.submit}
          color='primary'
          type='submit'
          variant='contained'
        >
          Join an existing game
        </Button>
      </Paper>
    </div>
  );
};

export default JoinGame;
