import { JoinGameState, JoinGameAction } from './types';
import { Reducer } from 'redux';

const initial_join_state: JoinGameState = {
  joinStatus: false,
  joinGameID: ''
};

export const joinReducer: Reducer<JoinGameState> = (
  state = initial_join_state,
  action
) => {
  return state;
};
