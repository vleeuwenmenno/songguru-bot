import ResourceState from "../types"

export interface Preferences {
    member: Preference
    guild: Preference
    guildId: string
}

export interface Preference {
    simpleMode: boolean                       
	mentionOnlyMode: boolean
	keepOriginalMessage: boolean

	allowOverrideSimpleMode: boolean|undefined
	allowOverrideMentionOnlyMode: boolean|undefined
	allowOverrideKeepOriginalMessage: boolean|undefined
}


export interface PreferencesState extends ResourceState<Preferences[]> {
}