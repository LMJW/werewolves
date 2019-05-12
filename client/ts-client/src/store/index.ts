import { combineReducers, Dispatch, Action, AnyAction } from 'redux';
import { loginReducer } from './login/reducer';
import { UserLoginState } from './login/types';
import { joinReducer } from './joinGame/reducer';
import { JoinGameState } from './joinGame/types';

export interface AppState {
  login: UserLoginState;
  join: JoinGameState;
}

export interface ConnectedReduxProps<A extends Action = AnyAction> {
  dispatch: Dispatch<A>;
}

export const RootReducer = combineReducers({
  login: loginReducer,
  join: joinReducer
});
