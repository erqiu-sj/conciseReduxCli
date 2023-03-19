import { CreateReducer , getAllValsWithActionCollectionHelper } from '@zealforchange/conciseredux'
 

 
 export type  actionAction = {
 ACTIONS: 'actions',
 PATAS: 'patas'
}
 
 export type actionActionTypes =  getAllValsWithActionCollectionHelper<actionAction>
 
 export type  actionStateTypes  = {}
 
 export type  actionActionPayloadTypes = {}
 
 export const action = new CreateReducer<actionStateTypes,actionActionPayloadTypes,actionActionTypes>({}) 
  .addAction('actions', (state, action) => {
    return { ...state }
  }) .addAction('patas', (state, action) => {
    return { ...state }
  }) 
 .setReducerKey('action')

