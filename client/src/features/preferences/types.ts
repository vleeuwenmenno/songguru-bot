import ResourceState from "../types"

export interface Preferences {
	guild: Guild
	guildSettings: GuildSettings
	memberSetting: any
}

export interface Guild {
	id: string
	name: string
	icon: string
	owner: boolean
	permissions: number
	permissions_new: string
	features: any[]
}

export interface GuildSettings {
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: any
	ID: string
	SimpleMode: boolean
	AllowOverrideSimpleMode: boolean
	MentionOnlyMode: boolean
	AllowOverrideMentionOnlyMode: boolean
	KeepOriginalMessage: boolean
	AllowOverrideKeepOriginalMessage: boolean
	GuildRefer: string
	Guild: GuildPreferences
}

export interface GuildPreferences {
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: any
	ID: string
	AdminRoleID: string
	JoinedAt: string
}

export interface PreferencesState extends ResourceState<Preferences[]> {
}