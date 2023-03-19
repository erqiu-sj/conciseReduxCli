import {action} from 'src/reducer/action';
import {like} from 'src/reducer/like';
import {combineReducers, createStore} from '@zealforchange/conciseredux'

const store = createStore(combineReducers({
    action: action.finish(),
    like: like.finish(),
}))

export {store}
