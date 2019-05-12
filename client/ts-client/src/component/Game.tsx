import React from 'react';
import useStyles from '../style/styles';
import { Typography, Paper } from '@material-ui/core';

interface GameProps {}

const Game = (props: GameProps) => {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <Paper className={classes.paper}>
        <h2>Room ID</h2>
        <Typography>Welcome</Typography>
      </Paper>
    </div>
  );
};

export default Game;
