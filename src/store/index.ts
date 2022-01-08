/*
 * @Author       : 邱狮杰
 * @Date         : 2021-07-29 14:03:55
 * @LastEditTime : 2021-09-24 09:44:25
 * @FilePath     : /preselectedweddingdress/src/reducers/index.ts
 * @Description  :
 */

import thunk from 'redux-thunk'
// import { createLogger } from 'redux-logger'
import {
    Common,
    FormOptionsInit,
    HeaderButtonReducer,
    HeaderCustomer,
    HeaderGoodsReducer,
    LikeAndDeleteIdCollection,
    Search,
    YarnSelectionListStatus
} from './modules'
import {applyMiddleware, combineReducers, createStore, Middleware} from '@zealforchange/conciseredux'

type behavioralConsumptionTimeOption = {
    log: string
    stop?: boolean
}
const behavioralConsumptionTime = function (option?: behavioralConsumptionTimeOption): Middleware {
    return () => next => action => {
        const pevTiem = new Date().getTime()
        const result = next(action)
        const nextTime = new Date().getTime()
        !option?.stop &&
        // @ts-ignore
        console[option?.log || 'group'](action.type, nextTime - pevTiem, 'ms')
        return result
    }
}
const middleware: any = [behavioralConsumptionTime({log: 'group', stop: true}), thunk,]
const store = createStore(
    combineReducers({
        HeaderButtonReducer: HeaderButtonReducer.finish(),
        HeaderGoodsReducer: HeaderGoodsReducer.finish(),
        HeaderCustomer: HeaderCustomer.finish(),
        YarnSelectionListStatus: YarnSelectionListStatus.finish(),
        Common: Common.finish(),
        FormOptionsInit: FormOptionsInit.finish(),
        LikeAndDeleteIdCollection: LikeAndDeleteIdCollection.finish(),
        Search: Search.finish(),
    }),
    applyMiddleware(...middleware)
)

export {store}
