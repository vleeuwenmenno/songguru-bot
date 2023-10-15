import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import { RootState } from "../../store"
import { Preferences, PreferencesState } from "./types"

const initialState: PreferencesState = {
    loading: false,
    data: [],
    error: '',
}

export const fetchPreferences = createAsyncThunk('preferences/fetch', async (_, { getState }) => {
    const state = getState() as RootState
    if (!state.auth.isAuthenticated) {
        throw new Error('User is not authenticated') // Throw an error if the user is not authenticated
    }

    const url = 'http://localhost:8081/api/settings'
    const options = { method: 'GET', credentials: 'include' } as RequestInit

    const response = await fetch(url, options)
    if (response.ok) {
        const json = await response.json()
        return json as Preferences[]
    } else {
        throw new Error('Failed to fetch preferences')
    }
})

const authenticationSlice = createSlice({
    name: "auth",
    initialState,
    extraReducers: (builder) => {
        builder.addCase(fetchPreferences.pending, (state) => {
            state.loading = true
        })

        builder.addCase(fetchPreferences.rejected, (state, action) => {
            state.loading = false
            state.data = []
            state.error = action.error.message ?? ''
        })

        builder.addCase(fetchPreferences.fulfilled, (state, action) => {
            state.loading = false
            state.data = action.payload
            state.error = ''
        })
    },
    reducers: {}
})

export default authenticationSlice.reducer