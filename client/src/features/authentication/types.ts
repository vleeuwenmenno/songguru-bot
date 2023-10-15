import ResourceState from "../types"

export interface DiscordUser {
    id: string
    username: string
    avatar: string
    discriminator: string
    public_flags: number
    flags: number
    banner: string
    accent_color: any
    global_name: string
    avatar_decoration_data: any
    banner_color: any
    mfa_enabled: boolean
    locale: string
    premium_type: number
}


export interface AuthenticationState extends ResourceState<DiscordUser> {
    isAuthenticated: boolean
}

