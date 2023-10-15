import { createAsyncThunk, createSlice, PayloadAction } from "@reduxjs/toolkit"
import { AuthenticationState, DiscordUser } from "./types"

const initialState: AuthenticationState = {
    loading: false,
    isAuthenticated: false,
    data: {} as DiscordUser,
    error: '',
}

export const fetchWhoAmI = createAsyncThunk('auth/whoami', () => {
    const url = 'http://localhost:8081/api/auth/whoami'
    const options = {method: 'GET', credentials: "include"} as RequestInit
    
    return fetch(url, options).then(async (response)  => {
        const json = await response.json()
        return json as DiscordUser
    })
})

export const doLogout = createAsyncThunk('auth/logout', () => {
    const url = 'http://localhost:8081/api/auth/logout'
    const options = {method: 'GET', credentials: "include"} as RequestInit
    
    return fetch(url, options).then(async (response)  => {
        return response.status == 200
    })
})

const authenticationSlice = createSlice({
    name: "auth",
    initialState,
    extraReducers: (builder) => {
        builder.addCase(fetchWhoAmI.pending, (state) => {
            state.loading = true
            state.isAuthenticated = false
        })

        builder.addCase(fetchWhoAmI.rejected, (state, action) => {
            state.loading = false
            state.isAuthenticated = false
            state.data = {} as DiscordUser
            state.error = action.error.message ?? ''
        })

        builder.addCase(fetchWhoAmI.fulfilled, (state, action) => {
            state.loading = false
            state.data = action.payload
            state.error = ''
            state.isAuthenticated = true
        })
        
        builder.addCase(doLogout.pending, (state) => {
            state.loading = true
            state.isAuthenticated = false
        })

        builder.addCase(doLogout.rejected, (state, action) => {
            state.loading = false
            state.isAuthenticated = false
            state.data = {} as DiscordUser
            state.error = action.error.message ?? ''
        })

        builder.addCase(doLogout.fulfilled, (state, action) => {
            state.loading = false
            state.data = {} as DiscordUser
            state.error = ''
            state.isAuthenticated = false
        })
    },
    reducers: {}
})

export default authenticationSlice.reducer