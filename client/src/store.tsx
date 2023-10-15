import { configureStore } from '@reduxjs/toolkit'
import authenticationSlice from './features/authentication/authenticationSlice'
import preferencesSlice from './features/preferences/preferencesSlice'

export const store = configureStore({
    reducer: {
        auth: authenticationSlice,
        preferences: preferencesSlice,
    },
})

export type AppDispatch = typeof store.dispatch
export type RootState = ReturnType<typeof store.getState>
