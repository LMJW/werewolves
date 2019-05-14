import React from 'react';
import useStyles from '../style/styles';
import { Typography, Paper } from '@material-ui/core';
import { withRouter } from 'react-router-dom';

const Game = props => {
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

export default withRouter(Game);
