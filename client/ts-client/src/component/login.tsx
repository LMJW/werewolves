import * as React from 'react';
import {
  FormControl,
  Paper,
  InputLabel,
  Input,
  InputAdornment,
  Button,
  Typography
} from '@material-ui/core';
import { Theme, createStyles, makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      display: 'flex',
      justifyContent: 'center'
    },
    paper: theme.mixins.gutters({
      paddingTop: 16,
      paddingBottom: 16,
      marginTop: theme.spacing(5),
      width: '30%',
      display: 'flex',
      flexDirection: 'column',
      alignContent: 'center',
      [theme.breakpoints.down('md')]: {
        width: '100%'
      }
    }),
    field: {
      marginTop: theme.spacing(5)
    },
    actions: theme.mixins.gutters({
      paddingTop: 16,
      paddingBottom: 16,
      marginTop: theme.spacing(5),
      display: 'flex',
      flexDirection: 'row',
      alignContent: 'center'
    }),
    submit: {
      marginTop: theme.spacing(5)
    }
  })
);

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
