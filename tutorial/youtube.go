package tutorial

type T struct {
	Auth struct {
		Oauth2 struct {
			Scopes struct {
				HttpsWwwGoogleapisComAuthYoutube struct {
					Description string `json:"description"`
				} `json:"https://www.googleapis.com/auth/youtube"`
				HttpsWwwGoogleapisComAuthYoutubeChannelMembershipsCreator struct {
					Description string `json:"description"`
				} `json:"https://www.googleapis.com/auth/youtube.channel-memberships.creator"`
				HttpsWwwGoogleapisComAuthYoutubeForceSsl struct {
					Description string `json:"description"`
				} `json:"https://www.googleapis.com/auth/youtube.force-ssl"`
				HttpsWwwGoogleapisComAuthYoutubeReadonly struct {
					Description string `json:"description"`
				} `json:"https://www.googleapis.com/auth/youtube.readonly"`
				HttpsWwwGoogleapisComAuthYoutubeUpload struct {
					Description string `json:"description"`
				} `json:"https://www.googleapis.com/auth/youtube.upload"`
				HttpsWwwGoogleapisComAuthYoutubepartner struct {
					Description string `json:"description"`
				} `json:"https://www.googleapis.com/auth/youtubepartner"`
				HttpsWwwGoogleapisComAuthYoutubepartnerChannelAudit struct {
					Description string `json:"description"`
				} `json:"https://www.googleapis.com/auth/youtubepartner-channel-audit"`
			} `json:"scopes"`
		} `json:"oauth2"`
	} `json:"auth"`
	BasePath                     string `json:"basePath"`
	BaseUrl                      string `json:"baseUrl"`
	BatchPath                    string `json:"batchPath"`
	CanonicalName                string `json:"canonicalName"`
	Description                  string `json:"description"`
	DiscoveryVersion             string `json:"discoveryVersion"`
	DocumentationLink            string `json:"documentationLink"`
	FullyEncodeReservedExpansion bool   `json:"fullyEncodeReservedExpansion"`
	Icons                        struct {
		X16 string `json:"x16"`
		X32 string `json:"x32"`
	} `json:"icons"`
	Id          string `json:"id"`
	Kind        string `json:"kind"`
	MtlsRootUrl string `json:"mtlsRootUrl"`
	Name        string `json:"name"`
	OwnerDomain string `json:"ownerDomain"`
	OwnerName   string `json:"ownerName"`
	Parameters  struct {
		Xgafv struct {
			Description      string   `json:"description"`
			Enum             []string `json:"enum"`
			EnumDescriptions []string `json:"enumDescriptions"`
			Location         string   `json:"location"`
			Type             string   `json:"type"`
		} `json:"$.xgafv"`
		AccessToken struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"access_token"`
		Alt struct {
			Default          string   `json:"default"`
			Description      string   `json:"description"`
			Enum             []string `json:"enum"`
			EnumDescriptions []string `json:"enumDescriptions"`
			Location         string   `json:"location"`
			Type             string   `json:"type"`
		} `json:"alt"`
		Callback struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"callback"`
		Fields struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"fields"`
		Key struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"key"`
		OauthToken struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"oauth_token"`
		PrettyPrint struct {
			Default     string `json:"default"`
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"prettyPrint"`
		QuotaUser struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"quotaUser"`
		UploadType struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"uploadType"`
		UploadProtocol struct {
			Description string `json:"description"`
			Location    string `json:"location"`
			Type        string `json:"type"`
		} `json:"upload_protocol"`
	} `json:"parameters"`
	Protocol  string `json:"protocol"`
	Resources struct {
		AbuseReports struct {
			Methods struct {
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
			} `json:"methods"`
		} `json:"abuseReports"`
		Activities struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"channelId"`
						Home struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"home"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						Mine struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"mine"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						PublishedAfter struct {
							Format   string `json:"format"`
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"publishedAfter"`
						PublishedBefore struct {
							Format   string `json:"format"`
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"publishedBefore"`
						RegionCode struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"regionCode"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"activities"`
		Captions struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						OnBehalfOf struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOf"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Download struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"id"`
						OnBehalfOf struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOf"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Tfmt struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"tfmt"`
						Tlang struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"tlang"`
					} `json:"parameters"`
					Path                    string   `json:"path"`
					Scopes                  []string `json:"scopes"`
					SupportsMediaDownload   bool     `json:"supportsMediaDownload"`
					UseMediaDownloadService bool     `json:"useMediaDownloadService"`
				} `json:"download"`
				Insert struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOf struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOf"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						Sync struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"sync"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						OnBehalfOf struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOf"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						VideoId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"videoId"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Update struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOf struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOf"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						Sync struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"sync"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"captions"`
		ChannelBanners struct {
			Methods struct {
				Insert struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"channelId"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"insert"`
			} `json:"methods"`
		} `json:"channelBanners"`
		ChannelSections struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"channelId"`
						Hl struct {
							Deprecated  bool   `json:"deprecated"`
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"hl"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						Mine struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"mine"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"channelSections"`
		Channels struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						CategoryId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"categoryId"`
						ForHandle struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"forHandle"`
						ForUsername struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"forUsername"`
						Hl struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"hl"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						ManagedByMe struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"managedByMe"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						Mine struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"mine"`
						MySubscribers struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"mySubscribers"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"channels"`
		CommentThreads struct {
			Methods struct {
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						AllThreadsRelatedToChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"allThreadsRelatedToChannelId"`
						ChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"channelId"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						ModerationStatus struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"moderationStatus"`
						Order struct {
							Default          string   `json:"default"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"order"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						SearchTerms struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"searchTerms"`
						TextFormat struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"textFormat"`
						VideoId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"videoId"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"commentThreads"`
		Comments struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						ParentId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"parentId"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						TextFormat struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"textFormat"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				MarkAsSpam struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"id"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"markAsSpam"`
				SetModerationStatus struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						BanAuthor struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"banAuthor"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"id"`
						ModerationStatus struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Required         bool     `json:"required"`
							Type             string   `json:"type"`
						} `json:"moderationStatus"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"setModerationStatus"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"comments"`
		I18NLanguages struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Hl struct {
							Default  string `json:"default"`
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"hl"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"i18nLanguages"`
		I18NRegions struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Hl struct {
							Default  string `json:"default"`
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"hl"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"i18nRegions"`
		LiveBroadcasts struct {
			Methods struct {
				Bind struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						StreamId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"streamId"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"bind"`
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				InsertCuepoint struct {
					Description    string        `json:"description"`
					FlatPath       string        `json:"flatPath"`
					HttpMethod     string        `json:"httpMethod"`
					Id             string        `json:"id"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insertCuepoint"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						BroadcastStatus struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"broadcastStatus"`
						BroadcastType struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"broadcastType"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						Mine struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"mine"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Transition struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						BroadcastStatus struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Required         bool     `json:"required"`
							Type             string   `json:"type"`
						} `json:"broadcastStatus"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"transition"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"liveBroadcasts"`
		LiveChatBans struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
			} `json:"methods"`
		} `json:"liveChatBans"`
		LiveChatMessages struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Hl struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"hl"`
						LiveChatId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"liveChatId"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						ProfileImageSize struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"profileImageSize"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Transition struct {
					Description    string        `json:"description"`
					FlatPath       string        `json:"flatPath"`
					HttpMethod     string        `json:"httpMethod"`
					Id             string        `json:"id"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"id"`
						Status struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"status"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"transition"`
			} `json:"methods"`
		} `json:"liveChatMessages"`
		LiveChatModerators struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						LiveChatId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"liveChatId"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"liveChatModerators"`
		LiveStreams struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						Mine struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"mine"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"liveStreams"`
		Members struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						FilterByMemberChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"filterByMemberChannelId"`
						HasAccessToLevel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"hasAccessToLevel"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						Mode struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"mode"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"members"`
		MembershipsLevels struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"membershipsLevels"`
		PlaylistImages struct {
			Methods struct {
				Delete struct {
					Description    string        `json:"description"`
					FlatPath       string        `json:"flatPath"`
					HttpMethod     string        `json:"httpMethod"`
					Id             string        `json:"id"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"insert"`
				List struct {
					Description    string        `json:"description"`
					FlatPath       string        `json:"flatPath"`
					HttpMethod     string        `json:"httpMethod"`
					Id             string        `json:"id"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Parent struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"parent"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Update struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"playlistImages"`
		PlaylistItems struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Repeated bool   `json:"repeated"`
							Type     string `json:"type"`
						} `json:"id"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						PlaylistId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"playlistId"`
						VideoId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"videoId"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"playlistItems"`
		Playlists struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"channelId"`
						Hl struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"hl"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						Mine struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"mine"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"playlists"`
		Search struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"channelId"`
						ChannelType struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"channelType"`
						EventType struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"eventType"`
						ForContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"forContentOwner"`
						ForDeveloper struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"forDeveloper"`
						ForMine struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"forMine"`
						Location struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"location"`
						LocationRadius struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"locationRadius"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Order struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"order"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						PublishedAfter struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"publishedAfter"`
						PublishedBefore struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"publishedBefore"`
						Q struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"q"`
						RegionCode struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"regionCode"`
						RelevanceLanguage struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"relevanceLanguage"`
						SafeSearch struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"safeSearch"`
						TopicId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"topicId"`
						Type struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"type"`
						VideoCaption struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoCaption"`
						VideoCategoryId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"videoCategoryId"`
						VideoDefinition struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoDefinition"`
						VideoDimension struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoDimension"`
						VideoDuration struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoDuration"`
						VideoEmbeddable struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoEmbeddable"`
						VideoLicense struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoLicense"`
						VideoPaidProductPlacement struct {
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoPaidProductPlacement"`
						VideoSyndicated struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoSyndicated"`
						VideoType struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"videoType"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"search"`
		Subscriptions struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"channelId"`
						ForChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"forChannelId"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						Mine struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"mine"`
						MyRecentSubscribers struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"myRecentSubscribers"`
						MySubscribers struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"mySubscribers"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Order struct {
							Default          string   `json:"default"`
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"order"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"subscriptions"`
		SuperChatEvents struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Hl struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"hl"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"superChatEvents"`
		Tests struct {
			Methods struct {
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ExternalChannelId struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"externalChannelId"`
						Part struct {
							Location string `json:"location"`
							Repeated bool   `json:"repeated"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"insert"`
			} `json:"methods"`
		} `json:"tests"`
		ThirdPartyLinks struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ExternalChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"externalChannelId"`
						LinkingToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"linkingToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"part"`
						Type struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Required         bool     `json:"required"`
							Type             string   `json:"type"`
						} `json:"type"`
					} `json:"parameters"`
					Path string `json:"path"`
				} `json:"delete"`
				Insert struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ExternalChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"externalChannelId"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ExternalChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"externalChannelId"`
						LinkingToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"linkingToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						Type struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"type"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
				} `json:"list"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ExternalChannelId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"externalChannelId"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"thirdPartyLinks"`
		Thumbnails struct {
			Methods struct {
				Set struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						VideoId struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"videoId"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"set"`
			} `json:"methods"`
		} `json:"thumbnails"`
		VideoAbuseReportReasons struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Hl struct {
							Default  string `json:"default"`
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"hl"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"videoAbuseReportReasons"`
		VideoCategories struct {
			Methods struct {
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Hl struct {
							Default  string `json:"default"`
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"hl"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						RegionCode struct {
							Location string `json:"location"`
							Type     string `json:"type"`
						} `json:"regionCode"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
			} `json:"methods"`
		} `json:"videoCategories"`
		Videos struct {
			Methods struct {
				Delete struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"delete"`
				GetRating struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Repeated bool   `json:"repeated"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"getRating"`
				Insert struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						AutoLevels struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"autoLevels"`
						NotifySubscribers struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"notifySubscribers"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						OnBehalfOfContentOwnerChannel struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwnerChannel"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						Stabilize struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"stabilize"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"insert"`
				List struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Chart struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"chart"`
						Hl struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"hl"`
						Id struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Type        string `json:"type"`
						} `json:"id"`
						Locale struct {
							Deprecated bool   `json:"deprecated"`
							Location   string `json:"location"`
							Type       string `json:"type"`
						} `json:"locale"`
						MaxHeight struct {
							Format   string `json:"format"`
							Location string `json:"location"`
							Maximum  string `json:"maximum"`
							Minimum  string `json:"minimum"`
							Type     string `json:"type"`
						} `json:"maxHeight"`
						MaxResults struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxResults"`
						MaxWidth struct {
							Description string `json:"description"`
							Format      string `json:"format"`
							Location    string `json:"location"`
							Maximum     string `json:"maximum"`
							Minimum     string `json:"minimum"`
							Type        string `json:"type"`
						} `json:"maxWidth"`
						MyRating struct {
							Description      string   `json:"description"`
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Type             string   `json:"type"`
						} `json:"myRating"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						PageToken struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"pageToken"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
						RegionCode struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"regionCode"`
						VideoCategoryId struct {
							Default     string `json:"default"`
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"videoCategoryId"`
					} `json:"parameters"`
					Path     string `json:"path"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"list"`
				Rate struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						Id struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"id"`
						Rating struct {
							Enum             []string `json:"enum"`
							EnumDescriptions []string `json:"enumDescriptions"`
							Location         string   `json:"location"`
							Required         bool     `json:"required"`
							Type             string   `json:"type"`
						} `json:"rating"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"rate"`
				ReportAbuse struct {
					Description    string        `json:"description"`
					FlatPath       string        `json:"flatPath"`
					HttpMethod     string        `json:"httpMethod"`
					Id             string        `json:"id"`
					ParameterOrder []interface{} `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Scopes []string `json:"scopes"`
				} `json:"reportAbuse"`
				Update struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
						Part struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Repeated    bool   `json:"repeated"`
							Required    bool   `json:"required"`
							Type        string `json:"type"`
						} `json:"part"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Response struct {
						Ref string `json:"$ref"`
					} `json:"response"`
					Scopes []string `json:"scopes"`
				} `json:"update"`
			} `json:"methods"`
		} `json:"videos"`
		Watermarks struct {
			Methods struct {
				Set struct {
					Description string `json:"description"`
					FlatPath    string `json:"flatPath"`
					HttpMethod  string `json:"httpMethod"`
					Id          string `json:"id"`
					MediaUpload struct {
						Accept    []string `json:"accept"`
						MaxSize   string   `json:"maxSize"`
						Protocols struct {
							Resumable struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"resumable"`
							Simple struct {
								Multipart bool   `json:"multipart"`
								Path      string `json:"path"`
							} `json:"simple"`
						} `json:"protocols"`
					} `json:"mediaUpload"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"channelId"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path    string `json:"path"`
					Request struct {
						Ref string `json:"$ref"`
					} `json:"request"`
					Scopes              []string `json:"scopes"`
					SupportsMediaUpload bool     `json:"supportsMediaUpload"`
				} `json:"set"`
				Unset struct {
					Description    string   `json:"description"`
					FlatPath       string   `json:"flatPath"`
					HttpMethod     string   `json:"httpMethod"`
					Id             string   `json:"id"`
					ParameterOrder []string `json:"parameterOrder"`
					Parameters     struct {
						ChannelId struct {
							Location string `json:"location"`
							Required bool   `json:"required"`
							Type     string `json:"type"`
						} `json:"channelId"`
						OnBehalfOfContentOwner struct {
							Description string `json:"description"`
							Location    string `json:"location"`
							Type        string `json:"type"`
						} `json:"onBehalfOfContentOwner"`
					} `json:"parameters"`
					Path   string   `json:"path"`
					Scopes []string `json:"scopes"`
				} `json:"unset"`
			} `json:"methods"`
		} `json:"watermarks"`
		Youtube struct {
			Resources struct {
				V3 struct {
					Methods struct {
						UpdateCommentThreads struct {
							Description    string        `json:"description"`
							FlatPath       string        `json:"flatPath"`
							HttpMethod     string        `json:"httpMethod"`
							Id             string        `json:"id"`
							ParameterOrder []interface{} `json:"parameterOrder"`
							Parameters     struct {
								Part struct {
									Description string `json:"description"`
									Location    string `json:"location"`
									Repeated    bool   `json:"repeated"`
									Type        string `json:"type"`
								} `json:"part"`
							} `json:"parameters"`
							Path    string `json:"path"`
							Request struct {
								Ref string `json:"$ref"`
							} `json:"request"`
							Response struct {
								Ref string `json:"$ref"`
							} `json:"response"`
						} `json:"updateCommentThreads"`
					} `json:"methods"`
				} `json:"v3"`
			} `json:"resources"`
		} `json:"youtube"`
	} `json:"resources"`
	Revision string `json:"revision"`
	RootUrl  string `json:"rootUrl"`
	Schemas  struct {
		AbuseReport struct {
			Id         string `json:"id"`
			Properties struct {
				AbuseTypes struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"abuseTypes"`
				Description struct {
					Type string `json:"type"`
				} `json:"description"`
				RelatedEntities struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"relatedEntities"`
				Subject struct {
					Ref string `json:"$ref"`
				} `json:"subject"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"AbuseReport"`
		AbuseType struct {
			Id         string `json:"id"`
			Properties struct {
				Id struct {
					Type string `json:"type"`
				} `json:"id"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"AbuseType"`
		AccessPolicy struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Allowed struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"allowed"`
				Exception struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"exception"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"AccessPolicy"`
		Activity struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Activity"`
		ActivityContentDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Bulletin struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"bulletin"`
				ChannelItem struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"channelItem"`
				Comment struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"comment"`
				Favorite struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"favorite"`
				Like struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"like"`
				PlaylistItem struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"playlistItem"`
				PromotedItem struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"promotedItem"`
				Recommendation struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"recommendation"`
				Social struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"social"`
				Subscription struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"subscription"`
				Upload struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"upload"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetails"`
		ActivityContentDetailsBulletin struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsBulletin"`
		ActivityContentDetailsChannelItem struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsChannelItem"`
		ActivityContentDetailsComment struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsComment"`
		ActivityContentDetailsFavorite struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsFavorite"`
		ActivityContentDetailsLike struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsLike"`
		ActivityContentDetailsPlaylistItem struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				PlaylistId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"playlistId"`
				PlaylistItemId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"playlistItemId"`
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsPlaylistItem"`
		ActivityContentDetailsPromotedItem struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AdTag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"adTag"`
				ClickTrackingUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"clickTrackingUrl"`
				CreativeViewUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"creativeViewUrl"`
				CtaType struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"ctaType"`
				CustomCtaButtonText struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"customCtaButtonText"`
				DescriptionText struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"descriptionText"`
				DestinationUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"destinationUrl"`
				ForecastingUrl struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"forecastingUrl"`
				ImpressionUrl struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"impressionUrl"`
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsPromotedItem"`
		ActivityContentDetailsRecommendation struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Reason struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"reason"`
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
				SeedResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"seedResourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsRecommendation"`
		ActivityContentDetailsSocial struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Author struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"author"`
				ImageUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"imageUrl"`
				ReferenceUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"referenceUrl"`
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsSocial"`
		ActivityContentDetailsSubscription struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ResourceId struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"resourceId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsSubscription"`
		ActivityContentDetailsUpload struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityContentDetailsUpload"`
		ActivityListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivityListResponse"`
		ActivitySnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelTitle struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelTitle"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				GroupId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"groupId"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ActivitySnippet"`
		Caption struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Caption"`
		CaptionListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CaptionListResponse"`
		CaptionSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AudioTrackType struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"audioTrackType"`
				FailureReason struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"failureReason"`
				IsAutoSynced struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isAutoSynced"`
				IsCC struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isCC"`
				IsDraft struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isDraft"`
				IsEasyReader struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isEasyReader"`
				IsLarge struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isLarge"`
				Language struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"language"`
				LastUpdated struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"lastUpdated"`
				Name struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"name"`
				Status struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"status"`
				TrackKind struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"trackKind"`
				VideoId struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CaptionSnippet"`
		CdnSettings struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Format struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"format"`
				FrameRate struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"frameRate"`
				IngestionInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"ingestionInfo"`
				IngestionType struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"ingestionType"`
				Resolution struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"resolution"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CdnSettings"`
		Channel struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AuditDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"auditDetails"`
				BrandingSettings struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"brandingSettings"`
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				ContentOwnerDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentOwnerDetails"`
				ConversionPings struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"conversionPings"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Localizations struct {
					AdditionalProperties struct {
						Ref string `json:"$ref"`
					} `json:"additionalProperties"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"localizations"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Statistics struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"statistics"`
				Status struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"status"`
				TopicDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"topicDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Channel"`
		ChannelAuditDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CommunityGuidelinesGoodStanding struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"communityGuidelinesGoodStanding"`
				ContentIdClaimsGoodStanding struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"contentIdClaimsGoodStanding"`
				CopyrightStrikesGoodStanding struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"copyrightStrikesGoodStanding"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelAuditDetails"`
		ChannelBannerResource struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Type string `json:"type"`
				} `json:"etag"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Url struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"url"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelBannerResource"`
		ChannelBrandingSettings struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Channel struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"channel"`
				Hints struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"hints"`
				Image struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"image"`
				Watch struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"watch"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelBrandingSettings"`
		ChannelContentDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				RelatedPlaylists struct {
					Properties struct {
						Favorites struct {
							Deprecated  bool   `json:"deprecated"`
							Description string `json:"description"`
							Type        string `json:"type"`
						} `json:"favorites"`
						Likes struct {
							Description string `json:"description"`
							Type        string `json:"type"`
						} `json:"likes"`
						Uploads struct {
							Description string `json:"description"`
							Type        string `json:"type"`
						} `json:"uploads"`
						WatchHistory struct {
							Deprecated  bool   `json:"deprecated"`
							Description string `json:"description"`
							Type        string `json:"type"`
						} `json:"watchHistory"`
						WatchLater struct {
							Deprecated  bool   `json:"deprecated"`
							Description string `json:"description"`
							Type        string `json:"type"`
						} `json:"watchLater"`
					} `json:"properties"`
					Type string `json:"type"`
				} `json:"relatedPlaylists"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelContentDetails"`
		ChannelContentOwnerDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ContentOwner struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"contentOwner"`
				TimeLinked struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"timeLinked"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelContentOwnerDetails"`
		ChannelConversionPing struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Context struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"context"`
				ConversionUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"conversionUrl"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelConversionPing"`
		ChannelConversionPings struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Pings struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"pings"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelConversionPings"`
		ChannelListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelListResponse"`
		ChannelLocalization struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelLocalization"`
		ChannelProfileDetails struct {
			Id         string `json:"id"`
			Properties struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelUrl"`
				DisplayName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"displayName"`
				ProfileImageUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"profileImageUrl"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelProfileDetails"`
		ChannelSection struct {
			Id         string `json:"id"`
			Properties struct {
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Localizations struct {
					AdditionalProperties struct {
						Ref string `json:"$ref"`
					} `json:"additionalProperties"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"localizations"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Targeting struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"targeting"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSection"`
		ChannelSectionContentDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Channels struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"channels"`
				Playlists struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"playlists"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSectionContentDetails"`
		ChannelSectionListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSectionListResponse"`
		ChannelSectionLocalization struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Title struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSectionLocalization"`
		ChannelSectionSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				DefaultLanguage struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"defaultLanguage"`
				Localized struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"localized"`
				Position struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"position"`
				Style struct {
					Deprecated       bool     `json:"deprecated"`
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"style"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDeprecated   []bool   `json:"enumDeprecated"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSectionSnippet"`
		ChannelSectionTargeting struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Countries struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"countries"`
				Languages struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"languages"`
				Regions struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"regions"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSectionTargeting"`
		ChannelSettings struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Country struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"country"`
				DefaultLanguage struct {
					Type string `json:"type"`
				} `json:"defaultLanguage"`
				DefaultTab struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"defaultTab"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				FeaturedChannelsTitle struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"featuredChannelsTitle"`
				FeaturedChannelsUrls struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"featuredChannelsUrls"`
				Keywords struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"keywords"`
				ModerateComments struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"moderateComments"`
				ProfileColor struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"profileColor"`
				ShowBrowseView struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"showBrowseView"`
				ShowRelatedChannels struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"showRelatedChannels"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
				TrackingAnalyticsAccountId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"trackingAnalyticsAccountId"`
				UnsubscribedTrailer struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"unsubscribedTrailer"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSettings"`
		ChannelSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Country struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"country"`
				CustomUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"customUrl"`
				DefaultLanguage struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"defaultLanguage"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				Localized struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"localized"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelSnippet"`
		ChannelStatistics struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CommentCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"commentCount"`
				HiddenSubscriberCount struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"hiddenSubscriberCount"`
				SubscriberCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"subscriberCount"`
				VideoCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"videoCount"`
				ViewCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"viewCount"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelStatistics"`
		ChannelStatus struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				IsLinked struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isLinked"`
				LongUploadsStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"longUploadsStatus"`
				MadeForKids struct {
					Type string `json:"type"`
				} `json:"madeForKids"`
				PrivacyStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"privacyStatus"`
				SelfDeclaredMadeForKids struct {
					Type string `json:"type"`
				} `json:"selfDeclaredMadeForKids"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelStatus"`
		ChannelToStoreLinkDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BillingDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"billingDetails"`
				MerchantId struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"merchantId"`
				StoreName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"storeName"`
				StoreUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"storeUrl"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelToStoreLinkDetails"`
		ChannelToStoreLinkDetailsBillingDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BillingStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"billingStatus"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelToStoreLinkDetailsBillingDetails"`
		ChannelTopicDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				TopicCategories struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"topicCategories"`
				TopicIds struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"topicIds"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ChannelTopicDetails"`
		Comment struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Comment"`
		CommentListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CommentListResponse"`
		CommentSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AuthorChannelId struct {
					Ref string `json:"$ref"`
				} `json:"authorChannelId"`
				AuthorChannelUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"authorChannelUrl"`
				AuthorDisplayName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"authorDisplayName"`
				AuthorProfileImageUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"authorProfileImageUrl"`
				CanRate struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"canRate"`
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				LikeCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"likeCount"`
				ModerationStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"moderationStatus"`
				ParentId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"parentId"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				TextDisplay struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"textDisplay"`
				TextOriginal struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"textOriginal"`
				UpdatedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"updatedAt"`
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
				ViewerRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"viewerRating"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CommentSnippet"`
		CommentSnippetAuthorChannelId struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Value struct {
					Type string `json:"type"`
				} `json:"value"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CommentSnippetAuthorChannelId"`
		CommentThread struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Replies struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"replies"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CommentThread"`
		CommentThreadListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CommentThreadListResponse"`
		CommentThreadReplies struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Comments struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"comments"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CommentThreadReplies"`
		CommentThreadSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CanReply struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"canReply"`
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				IsPublic struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isPublic"`
				TopLevelComment struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"topLevelComment"`
				TotalReplyCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"totalReplyCount"`
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CommentThreadSnippet"`
		ContentRating struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AcbRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"acbRating"`
				AgcomRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"agcomRating"`
				AnatelRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"anatelRating"`
				BbfcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"bbfcRating"`
				BfvcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"bfvcRating"`
				BmukkRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"bmukkRating"`
				CatvRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"catvRating"`
				CatvfrRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"catvfrRating"`
				CbfcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cbfcRating"`
				CccRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cccRating"`
				CceRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cceRating"`
				ChfilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"chfilmRating"`
				ChvrsRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"chvrsRating"`
				CicfRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cicfRating"`
				CnaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cnaRating"`
				CncRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cncRating"`
				CsaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"csaRating"`
				CscfRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cscfRating"`
				CzfilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"czfilmRating"`
				DjctqRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"djctqRating"`
				DjctqRatingReasons struct {
					Description string `json:"description"`
					Items       struct {
						Enum             []string `json:"enum"`
						EnumDescriptions []string `json:"enumDescriptions"`
						Type             string   `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"djctqRatingReasons"`
				EcbmctRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"ecbmctRating"`
				EefilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"eefilmRating"`
				EgfilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"egfilmRating"`
				EirinRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"eirinRating"`
				FcbmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"fcbmRating"`
				FcoRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"fcoRating"`
				FmocRating struct {
					Deprecated       bool     `json:"deprecated"`
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"fmocRating"`
				FpbRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"fpbRating"`
				FpbRatingReasons struct {
					Description string `json:"description"`
					Items       struct {
						Enum             []string `json:"enum"`
						EnumDescriptions []string `json:"enumDescriptions"`
						Type             string   `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"fpbRatingReasons"`
				FskRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"fskRating"`
				GrfilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"grfilmRating"`
				IcaaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"icaaRating"`
				IfcoRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"ifcoRating"`
				IlfilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"ilfilmRating"`
				IncaaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"incaaRating"`
				KfcbRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"kfcbRating"`
				KijkwijzerRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"kijkwijzerRating"`
				KmrbRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"kmrbRating"`
				LsfRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDeprecated   []bool   `json:"enumDeprecated"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"lsfRating"`
				MccaaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mccaaRating"`
				MccypRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mccypRating"`
				McstRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mcstRating"`
				MdaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mdaRating"`
				MedietilsynetRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"medietilsynetRating"`
				MekuRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mekuRating"`
				MenaMpaaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"menaMpaaRating"`
				MibacRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mibacRating"`
				MocRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mocRating"`
				MoctwRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"moctwRating"`
				MpaaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mpaaRating"`
				MpaatRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mpaatRating"`
				MtrcbRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"mtrcbRating"`
				NbcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"nbcRating"`
				NbcplRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"nbcplRating"`
				NfrcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"nfrcRating"`
				NfvcbRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"nfvcbRating"`
				NkclvRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"nkclvRating"`
				NmcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"nmcRating"`
				OflcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"oflcRating"`
				PefilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"pefilmRating"`
				RcnofRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"rcnofRating"`
				ResorteviolenciaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"resorteviolenciaRating"`
				RtcRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"rtcRating"`
				RteRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"rteRating"`
				RussiaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"russiaRating"`
				SkfilmRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"skfilmRating"`
				SmaisRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"smaisRating"`
				SmsaRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"smsaRating"`
				TvpgRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"tvpgRating"`
				YtRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"ytRating"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ContentRating"`
		Cuepoint struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CueType struct {
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cueType"`
				DurationSecs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"durationSecs"`
				Etag struct {
					Type string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				InsertionOffsetTimeMs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"insertionOffsetTimeMs"`
				WalltimeMs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"walltimeMs"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Cuepoint"`
		CuepointSchedule struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Enabled struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enabled"`
				PauseAdsUntil struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"pauseAdsUntil"`
				RepeatIntervalSecs struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"repeatIntervalSecs"`
				ScheduleStrategy struct {
					Deprecated       bool     `json:"deprecated"`
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"scheduleStrategy"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"CuepointSchedule"`
		Entity struct {
			Id         string `json:"id"`
			Properties struct {
				Id struct {
					Type string `json:"type"`
				} `json:"id"`
				TypeId struct {
					Type string `json:"type"`
				} `json:"typeId"`
				Url struct {
					Type string `json:"type"`
				} `json:"url"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Entity"`
		GeoPoint struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Altitude struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"altitude"`
				Latitude struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"latitude"`
				Longitude struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"longitude"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"GeoPoint"`
		I18NLanguage struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"I18nLanguage"`
		I18NLanguageListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"I18nLanguageListResponse"`
		I18NLanguageSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Hl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"hl"`
				Name struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"name"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"I18nLanguageSnippet"`
		I18NRegion struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"I18nRegion"`
		I18NRegionListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"I18nRegionListResponse"`
		I18NRegionSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Gl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"gl"`
				Name struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"name"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"I18nRegionSnippet"`
		ImageSettings struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BackgroundImageUrl struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"backgroundImageUrl"`
				BannerExternalUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerExternalUrl"`
				BannerImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerImageUrl"`
				BannerMobileExtraHdImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerMobileExtraHdImageUrl"`
				BannerMobileHdImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerMobileHdImageUrl"`
				BannerMobileImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerMobileImageUrl"`
				BannerMobileLowImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerMobileLowImageUrl"`
				BannerMobileMediumHdImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerMobileMediumHdImageUrl"`
				BannerTabletExtraHdImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTabletExtraHdImageUrl"`
				BannerTabletHdImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTabletHdImageUrl"`
				BannerTabletImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTabletImageUrl"`
				BannerTabletLowImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTabletLowImageUrl"`
				BannerTvHighImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTvHighImageUrl"`
				BannerTvImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTvImageUrl"`
				BannerTvLowImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTvLowImageUrl"`
				BannerTvMediumImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"bannerTvMediumImageUrl"`
				LargeBrandedBannerImageImapScript struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"largeBrandedBannerImageImapScript"`
				LargeBrandedBannerImageUrl struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"largeBrandedBannerImageUrl"`
				SmallBrandedBannerImageImapScript struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"smallBrandedBannerImageImapScript"`
				SmallBrandedBannerImageUrl struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"smallBrandedBannerImageUrl"`
				TrackingImageUrl struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"trackingImageUrl"`
				WatchIconImageUrl struct {
					Deprecated bool   `json:"deprecated"`
					Type       string `json:"type"`
				} `json:"watchIconImageUrl"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ImageSettings"`
		IngestionInfo struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BackupIngestionAddress struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"backupIngestionAddress"`
				IngestionAddress struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"ingestionAddress"`
				RtmpsBackupIngestionAddress struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"rtmpsBackupIngestionAddress"`
				RtmpsIngestionAddress struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"rtmpsIngestionAddress"`
				StreamName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"streamName"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"IngestionInfo"`
		InvideoBranding struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ImageBytes struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"imageBytes"`
				ImageUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"imageUrl"`
				Position struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"position"`
				TargetChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"targetChannelId"`
				Timing struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"timing"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"InvideoBranding"`
		InvideoPosition struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CornerPosition struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDeprecated   []bool   `json:"enumDeprecated"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"cornerPosition"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"InvideoPosition"`
		InvideoTiming struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				DurationMs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"durationMs"`
				OffsetMs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"offsetMs"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"InvideoTiming"`
		LanguageTag struct {
			Id         string `json:"id"`
			Properties struct {
				Value struct {
					Type string `json:"type"`
				} `json:"value"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LanguageTag"`
		LevelDetails struct {
			Id         string `json:"id"`
			Properties struct {
				DisplayName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"displayName"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LevelDetails"`
		LiveBroadcast struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				MonetizationDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"monetizationDetails"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Statistics struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"statistics"`
				Status struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveBroadcast"`
		LiveBroadcastContentDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BoundStreamId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"boundStreamId"`
				BoundStreamLastUpdateTimeMs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"boundStreamLastUpdateTimeMs"`
				ClosedCaptionsType struct {
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"closedCaptionsType"`
				EnableAutoStart struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableAutoStart"`
				EnableAutoStop struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableAutoStop"`
				EnableClosedCaptions struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableClosedCaptions"`
				EnableContentEncryption struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableContentEncryption"`
				EnableDvr struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableDvr"`
				EnableEmbed struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableEmbed"`
				EnableLowLatency struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableLowLatency"`
				LatencyPreference struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"latencyPreference"`
				Mesh struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"mesh"`
				MonitorStream struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"monitorStream"`
				Projection struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"projection"`
				RecordFromStart struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"recordFromStart"`
				StartWithSlate struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"startWithSlate"`
				StereoLayout struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"stereoLayout"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveBroadcastContentDetails"`
		LiveBroadcastListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveBroadcastListResponse"`
		LiveBroadcastMonetizationDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CuepointSchedule struct {
					Ref string `json:"$ref"`
				} `json:"cuepointSchedule"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveBroadcastMonetizationDetails"`
		LiveBroadcastSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ActualEndTime struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"actualEndTime"`
				ActualStartTime struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"actualStartTime"`
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				IsDefaultBroadcast struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isDefaultBroadcast"`
				LiveChatId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"liveChatId"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				ScheduledEndTime struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"scheduledEndTime"`
				ScheduledStartTime struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"scheduledStartTime"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveBroadcastSnippet"`
		LiveBroadcastStatistics struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ConcurrentViewers struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"concurrentViewers"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveBroadcastStatistics"`
		LiveBroadcastStatus struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				LifeCycleStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"lifeCycleStatus"`
				LiveBroadcastPriority struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"liveBroadcastPriority"`
				MadeForKids struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"madeForKids"`
				PrivacyStatus struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"privacyStatus"`
				RecordingStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"recordingStatus"`
				SelfDeclaredMadeForKids struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"selfDeclaredMadeForKids"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveBroadcastStatus"`
		LiveChatBan struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatBan"`
		LiveChatBanSnippet struct {
			Id         string `json:"id"`
			Properties struct {
				BanDurationSeconds struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"banDurationSeconds"`
				BannedUserDetails struct {
					Ref string `json:"$ref"`
				} `json:"bannedUserDetails"`
				LiveChatId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"liveChatId"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatBanSnippet"`
		LiveChatFanFundingEventDetails struct {
			Id         string `json:"id"`
			Properties struct {
				AmountDisplayString struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"amountDisplayString"`
				AmountMicros struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"amountMicros"`
				Currency struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"currency"`
				UserComment struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"userComment"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatFanFundingEventDetails"`
		LiveChatGiftMembershipReceivedDetails struct {
			Id         string `json:"id"`
			Properties struct {
				AssociatedMembershipGiftingMessageId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"associatedMembershipGiftingMessageId"`
				GifterChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"gifterChannelId"`
				MemberLevelName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"memberLevelName"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatGiftMembershipReceivedDetails"`
		LiveChatMemberMilestoneChatDetails struct {
			Id         string `json:"id"`
			Properties struct {
				MemberLevelName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"memberLevelName"`
				MemberMonth struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"memberMonth"`
				UserComment struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"userComment"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMemberMilestoneChatDetails"`
		LiveChatMembershipGiftingDetails struct {
			Id         string `json:"id"`
			Properties struct {
				GiftMembershipsCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"giftMembershipsCount"`
				GiftMembershipsLevelName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"giftMembershipsLevelName"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMembershipGiftingDetails"`
		LiveChatMessage struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AuthorDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"authorDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMessage"`
		LiveChatMessageAuthorDetails struct {
			Id         string `json:"id"`
			Properties struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelUrl"`
				DisplayName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"displayName"`
				IsChatModerator struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isChatModerator"`
				IsChatOwner struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isChatOwner"`
				IsChatSponsor struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isChatSponsor"`
				IsVerified struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isVerified"`
				ProfileImageUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"profileImageUrl"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMessageAuthorDetails"`
		LiveChatMessageDeletedDetails struct {
			Id         string `json:"id"`
			Properties struct {
				DeletedMessageId struct {
					Type string `json:"type"`
				} `json:"deletedMessageId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMessageDeletedDetails"`
		LiveChatMessageListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				ActivePollItem struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"activePollItem"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Type string `json:"type"`
				} `json:"nextPageToken"`
				OfflineAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"offlineAt"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PollingIntervalMillis struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"pollingIntervalMillis"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMessageListResponse"`
		LiveChatMessageRetractedDetails struct {
			Id         string `json:"id"`
			Properties struct {
				RetractedMessageId struct {
					Type string `json:"type"`
				} `json:"retractedMessageId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMessageRetractedDetails"`
		LiveChatMessageSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AuthorChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"authorChannelId"`
				DisplayMessage struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"displayMessage"`
				FanFundingEventDetails struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"fanFundingEventDetails"`
				GiftMembershipReceivedDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"giftMembershipReceivedDetails"`
				HasDisplayContent struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"hasDisplayContent"`
				LiveChatId struct {
					Type string `json:"type"`
				} `json:"liveChatId"`
				MemberMilestoneChatDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"memberMilestoneChatDetails"`
				MembershipGiftingDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"membershipGiftingDetails"`
				MessageDeletedDetails struct {
					Ref string `json:"$ref"`
				} `json:"messageDeletedDetails"`
				MessageRetractedDetails struct {
					Ref string `json:"$ref"`
				} `json:"messageRetractedDetails"`
				NewSponsorDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"newSponsorDetails"`
				PollDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pollDetails"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				SuperChatDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"superChatDetails"`
				SuperStickerDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"superStickerDetails"`
				TextMessageDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"textMessageDetails"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
				UserBannedDetails struct {
					Ref string `json:"$ref"`
				} `json:"userBannedDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatMessageSnippet"`
		LiveChatModerator struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatModerator"`
		LiveChatModeratorListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatModeratorListResponse"`
		LiveChatModeratorSnippet struct {
			Id         string `json:"id"`
			Properties struct {
				LiveChatId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"liveChatId"`
				ModeratorDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"moderatorDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatModeratorSnippet"`
		LiveChatNewSponsorDetails struct {
			Id         string `json:"id"`
			Properties struct {
				IsUpgrade struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isUpgrade"`
				MemberLevelName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"memberLevelName"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatNewSponsorDetails"`
		LiveChatPollDetails struct {
			Id         string `json:"id"`
			Properties struct {
				Metadata struct {
					Ref string `json:"$ref"`
				} `json:"metadata"`
				Status struct {
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"status"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatPollDetails"`
		LiveChatPollDetailsPollMetadata struct {
			Id         string `json:"id"`
			Properties struct {
				Options struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"options"`
				QuestionText struct {
					Type string `json:"type"`
				} `json:"questionText"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatPollDetailsPollMetadata"`
		LiveChatPollDetailsPollMetadataPollOption struct {
			Id         string `json:"id"`
			Properties struct {
				OptionText struct {
					Type string `json:"type"`
				} `json:"optionText"`
				Tally struct {
					Format string `json:"format"`
					Type   string `json:"type"`
				} `json:"tally"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatPollDetailsPollMetadataPollOption"`
		LiveChatSuperChatDetails struct {
			Id         string `json:"id"`
			Properties struct {
				AmountDisplayString struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"amountDisplayString"`
				AmountMicros struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"amountMicros"`
				Currency struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"currency"`
				Tier struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"tier"`
				UserComment struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"userComment"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatSuperChatDetails"`
		LiveChatSuperStickerDetails struct {
			Id         string `json:"id"`
			Properties struct {
				AmountDisplayString struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"amountDisplayString"`
				AmountMicros struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"amountMicros"`
				Currency struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"currency"`
				SuperStickerMetadata struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"superStickerMetadata"`
				Tier struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"tier"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatSuperStickerDetails"`
		LiveChatTextMessageDetails struct {
			Id         string `json:"id"`
			Properties struct {
				MessageText struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"messageText"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatTextMessageDetails"`
		LiveChatUserBannedMessageDetails struct {
			Id         string `json:"id"`
			Properties struct {
				BanDurationSeconds struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"banDurationSeconds"`
				BanType struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"banType"`
				BannedUserDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"bannedUserDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveChatUserBannedMessageDetails"`
		LiveStream struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Cdn struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"cdn"`
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Status struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveStream"`
		LiveStreamConfigurationIssue struct {
			Id         string `json:"id"`
			Properties struct {
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				Reason struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"reason"`
				Severity struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"severity"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveStreamConfigurationIssue"`
		LiveStreamContentDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ClosedCaptionsIngestionUrl struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"closedCaptionsIngestionUrl"`
				IsReusable struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isReusable"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveStreamContentDetails"`
		LiveStreamHealthStatus struct {
			Id         string `json:"id"`
			Properties struct {
				ConfigurationIssues struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"configurationIssues"`
				LastUpdateTimeSeconds struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"lastUpdateTimeSeconds"`
				Status struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"status"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveStreamHealthStatus"`
		LiveStreamListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref string `json:"$ref"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveStreamListResponse"`
		LiveStreamSnippet struct {
			Id         string `json:"id"`
			Properties struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				IsDefaultStream struct {
					Type string `json:"type"`
				} `json:"isDefaultStream"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				Title struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveStreamSnippet"`
		LiveStreamStatus struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				HealthStatus struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"healthStatus"`
				StreamStatus struct {
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"streamStatus"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LiveStreamStatus"`
		LocalizedProperty struct {
			Id         string `json:"id"`
			Properties struct {
				Default struct {
					Type string `json:"type"`
				} `json:"default"`
				DefaultLanguage struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"defaultLanguage"`
				Localized struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"localized"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LocalizedProperty"`
		LocalizedString struct {
			Id         string `json:"id"`
			Properties struct {
				Language struct {
					Type string `json:"type"`
				} `json:"language"`
				Value struct {
					Type string `json:"type"`
				} `json:"value"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"LocalizedString"`
		Member struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Member"`
		MemberListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref string `json:"$ref"`
				} `json:"pageInfo"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MemberListResponse"`
		MemberSnippet struct {
			Id         string `json:"id"`
			Properties struct {
				CreatorChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"creatorChannelId"`
				MemberDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"memberDetails"`
				MembershipsDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"membershipsDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MemberSnippet"`
		MembershipsDetails struct {
			Id         string `json:"id"`
			Properties struct {
				AccessibleLevels struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"accessibleLevels"`
				HighestAccessibleLevel struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"highestAccessibleLevel"`
				HighestAccessibleLevelDisplayName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"highestAccessibleLevelDisplayName"`
				MembershipsDuration struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"membershipsDuration"`
				MembershipsDurationAtLevels struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"membershipsDurationAtLevels"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MembershipsDetails"`
		MembershipsDuration struct {
			Id         string `json:"id"`
			Properties struct {
				MemberSince struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"memberSince"`
				MemberTotalDurationMonths struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"memberTotalDurationMonths"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MembershipsDuration"`
		MembershipsDurationAtLevel struct {
			Id         string `json:"id"`
			Properties struct {
				Level struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"level"`
				MemberSince struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"memberSince"`
				MemberTotalDurationMonths struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"memberTotalDurationMonths"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MembershipsDurationAtLevel"`
		MembershipsLevel struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MembershipsLevel"`
		MembershipsLevelListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MembershipsLevelListResponse"`
		MembershipsLevelSnippet struct {
			Id         string `json:"id"`
			Properties struct {
				CreatorChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"creatorChannelId"`
				LevelDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"levelDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MembershipsLevelSnippet"`
		MonitorStreamInfo struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BroadcastStreamDelayMs struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"broadcastStreamDelayMs"`
				EmbedHtml struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"embedHtml"`
				EnableMonitorStream struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"enableMonitorStream"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"MonitorStreamInfo"`
		PageInfo struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ResultsPerPage struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"resultsPerPage"`
				TotalResults struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"totalResults"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PageInfo"`
		Playlist struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Localizations struct {
					AdditionalProperties struct {
						Ref string `json:"$ref"`
					} `json:"additionalProperties"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"localizations"`
				Player struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"player"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Status struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Playlist"`
		PlaylistContentDetails struct {
			Id         string `json:"id"`
			Properties struct {
				ItemCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"itemCount"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistContentDetails"`
		PlaylistImage struct {
			Id         string `json:"id"`
			Properties struct {
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref string `json:"$ref"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistImage"`
		PlaylistImageListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Items struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistImageListResponse"`
		PlaylistImageSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Height struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"height"`
				PlaylistId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"playlistId"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
				Width struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"width"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistImageSnippet"`
		PlaylistItem struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Status struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistItem"`
		PlaylistItemContentDetails struct {
			Id         string `json:"id"`
			Properties struct {
				EndAt struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"endAt"`
				Note struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"note"`
				StartAt struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"startAt"`
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
				VideoPublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"videoPublishedAt"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistItemContentDetails"`
		PlaylistItemListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Type string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref string `json:"$ref"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistItemListResponse"`
		PlaylistItemSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelTitle struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelTitle"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				PlaylistId struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"playlistId"`
				Position struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"position"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				ResourceId struct {
					Ref         string `json:"$ref"`
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
				} `json:"resourceId"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
				VideoOwnerChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoOwnerChannelId"`
				VideoOwnerChannelTitle struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoOwnerChannelTitle"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistItemSnippet"`
		PlaylistItemStatus struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				PrivacyStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"privacyStatus"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistItemStatus"`
		PlaylistListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistListResponse"`
		PlaylistLocalization struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistLocalization"`
		PlaylistPlayer struct {
			Id         string `json:"id"`
			Properties struct {
				EmbedHtml struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"embedHtml"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistPlayer"`
		PlaylistSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelTitle struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelTitle"`
				DefaultLanguage struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"defaultLanguage"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				Localized struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"localized"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				Tags struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"tags"`
				ThumbnailVideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"thumbnailVideoId"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistSnippet"`
		PlaylistStatus struct {
			Id         string `json:"id"`
			Properties struct {
				PrivacyStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"privacyStatus"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PlaylistStatus"`
		PropertyValue struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Property struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"property"`
				Value struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"value"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"PropertyValue"`
		RelatedEntity struct {
			Id         string `json:"id"`
			Properties struct {
				Entity struct {
					Ref string `json:"$ref"`
				} `json:"entity"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"RelatedEntity"`
		ResourceId struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				Kind struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				PlaylistId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"playlistId"`
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ResourceId"`
		SearchListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				RegionCode struct {
					Type string `json:"type"`
				} `json:"regionCode"`
				TokenPagination struct {
					Ref string `json:"$ref"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SearchListResponse"`
		SearchResult struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SearchResult"`
		SearchResultSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelTitle struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelTitle"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				LiveBroadcastContent struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"liveBroadcastContent"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SearchResultSnippet"`
		Subscription struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				SubscriberSnippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"subscriberSnippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Subscription"`
		SubscriptionContentDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ActivityType struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"activityType"`
				NewItemCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"newItemCount"`
				TotalItemCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"totalItemCount"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SubscriptionContentDetails"`
		SubscriptionListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref string `json:"$ref"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SubscriptionListResponse"`
		SubscriptionSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelTitle struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelTitle"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				ResourceId struct {
					Ref         string `json:"$ref"`
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
				} `json:"resourceId"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SubscriptionSnippet"`
		SubscriptionSubscriberSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SubscriptionSubscriberSnippet"`
		SuperChatEvent struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SuperChatEvent"`
		SuperChatEventListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref string `json:"$ref"`
				} `json:"pageInfo"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SuperChatEventListResponse"`
		SuperChatEventSnippet struct {
			Id         string `json:"id"`
			Properties struct {
				AmountMicros struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"amountMicros"`
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				CommentText struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"commentText"`
				CreatedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"createdAt"`
				Currency struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"currency"`
				DisplayString struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"displayString"`
				IsSuperStickerEvent struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"isSuperStickerEvent"`
				MessageType struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"messageType"`
				SuperStickerMetadata struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"superStickerMetadata"`
				SupporterDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"supporterDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SuperChatEventSnippet"`
		SuperStickerMetadata struct {
			Id         string `json:"id"`
			Properties struct {
				AltText struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"altText"`
				AltTextLanguage struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"altTextLanguage"`
				StickerId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"stickerId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"SuperStickerMetadata"`
		TestItem struct {
			Id         string `json:"id"`
			Properties struct {
				FeaturedPart struct {
					Type string `json:"type"`
				} `json:"featuredPart"`
				Gaia struct {
					Format string `json:"format"`
					Type   string `json:"type"`
				} `json:"gaia"`
				Id struct {
					Type string `json:"type"`
				} `json:"id"`
				Snippet struct {
					Ref string `json:"$ref"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"TestItem"`
		TestItemTestItemSnippet struct {
			Id         string `json:"id"`
			Properties struct {
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"TestItemTestItemSnippet"`
		ThirdPartyLink struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				LinkingToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"linkingToken"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Status struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"status"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ThirdPartyLink"`
		ThirdPartyLinkListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Items struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ThirdPartyLinkListResponse"`
		ThirdPartyLinkSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ChannelToStoreLink struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"channelToStoreLink"`
				Type struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"type"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ThirdPartyLinkSnippet"`
		ThirdPartyLinkStatus struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				LinkStatus struct {
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"linkStatus"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ThirdPartyLinkStatus"`
		Thumbnail struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Height struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"height"`
				Url struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"url"`
				Width struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"width"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Thumbnail"`
		ThumbnailDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Default struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"default"`
				High struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"high"`
				Maxres struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"maxres"`
				Medium struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"medium"`
				Standard struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"standard"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ThumbnailDetails"`
		ThumbnailSetResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"ThumbnailSetResponse"`
		TokenPagination struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"TokenPagination"`
		Video struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AgeGating struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"ageGating"`
				ContentDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentDetails"`
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				FileDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"fileDetails"`
				Id struct {
					Annotations struct {
						Required []string `json:"required"`
					} `json:"annotations"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				LiveStreamingDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"liveStreamingDetails"`
				Localizations struct {
					AdditionalProperties struct {
						Ref string `json:"$ref"`
					} `json:"additionalProperties"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"localizations"`
				MonetizationDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"monetizationDetails"`
				Player struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"player"`
				ProcessingDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"processingDetails"`
				ProjectDetails struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"projectDetails"`
				RecordingDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"recordingDetails"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
				Statistics struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"statistics"`
				Status struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"status"`
				Suggestions struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"suggestions"`
				TopicDetails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"topicDetails"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"Video"`
		VideoAbuseReport struct {
			Id         string `json:"id"`
			Properties struct {
				Comments struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"comments"`
				Language struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"language"`
				ReasonId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"reasonId"`
				SecondaryReasonId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"secondaryReasonId"`
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoAbuseReport"`
		VideoAbuseReportReason struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoAbuseReportReason"`
		VideoAbuseReportReasonListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoAbuseReportReasonListResponse"`
		VideoAbuseReportReasonSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Label struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"label"`
				SecondaryReasons struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"secondaryReasons"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoAbuseReportReasonSnippet"`
		VideoAbuseReportSecondaryReason struct {
			Id         string `json:"id"`
			Properties struct {
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Label struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"label"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoAbuseReportSecondaryReason"`
		VideoAgeGating struct {
			Id         string `json:"id"`
			Properties struct {
				AlcoholContent struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"alcoholContent"`
				Restricted struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"restricted"`
				VideoGameRating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"videoGameRating"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoAgeGating"`
		VideoCategory struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				Id struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"id"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				Snippet struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"snippet"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoCategory"`
		VideoCategoryListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoCategoryListResponse"`
		VideoCategorySnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Assignable struct {
					Type string `json:"type"`
				} `json:"assignable"`
				ChannelId struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoCategorySnippet"`
		VideoContentDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Caption struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"caption"`
				ContentRating struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"contentRating"`
				CountryRestriction struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"countryRestriction"`
				Definition struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"definition"`
				Dimension struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"dimension"`
				Duration struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"duration"`
				HasCustomThumbnail struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"hasCustomThumbnail"`
				LicensedContent struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"licensedContent"`
				Projection struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"projection"`
				RegionRestriction struct {
					Ref         string `json:"$ref"`
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
				} `json:"regionRestriction"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoContentDetails"`
		VideoContentDetailsRegionRestriction struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Allowed struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"allowed"`
				Blocked struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"blocked"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoContentDetailsRegionRestriction"`
		VideoFileDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AudioStreams struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"audioStreams"`
				BitrateBps struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"bitrateBps"`
				Container struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"container"`
				CreationTime struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"creationTime"`
				DurationMs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"durationMs"`
				FileName struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"fileName"`
				FileSize struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"fileSize"`
				FileType struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"fileType"`
				VideoStreams struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"videoStreams"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoFileDetails"`
		VideoFileDetailsAudioStream struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BitrateBps struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"bitrateBps"`
				ChannelCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"channelCount"`
				Codec struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"codec"`
				Vendor struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"vendor"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoFileDetailsAudioStream"`
		VideoFileDetailsVideoStream struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				AspectRatio struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"aspectRatio"`
				BitrateBps struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"bitrateBps"`
				Codec struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"codec"`
				FrameRateFps struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"frameRateFps"`
				HeightPixels struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"heightPixels"`
				Rotation struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"rotation"`
				Vendor struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"vendor"`
				WidthPixels struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"widthPixels"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoFileDetailsVideoStream"`
		VideoGetRatingResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoGetRatingResponse"`
		VideoListResponse struct {
			Id         string `json:"id"`
			Properties struct {
				Etag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"etag"`
				EventId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"eventId"`
				Items struct {
					Items struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"items"`
				Kind struct {
					Default     string `json:"default"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"kind"`
				NextPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"nextPageToken"`
				PageInfo struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"pageInfo"`
				PrevPageToken struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"prevPageToken"`
				TokenPagination struct {
					Ref        string `json:"$ref"`
					Deprecated bool   `json:"deprecated"`
				} `json:"tokenPagination"`
				VisitorId struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"visitorId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoListResponse"`
		VideoLiveStreamingDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				ActiveLiveChatId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"activeLiveChatId"`
				ActualEndTime struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"actualEndTime"`
				ActualStartTime struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"actualStartTime"`
				ConcurrentViewers struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"concurrentViewers"`
				ScheduledEndTime struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"scheduledEndTime"`
				ScheduledStartTime struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"scheduledStartTime"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoLiveStreamingDetails"`
		VideoLocalization struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoLocalization"`
		VideoMonetizationDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Access struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"access"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoMonetizationDetails"`
		VideoPlayer struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				EmbedHeight struct {
					Format string `json:"format"`
					Type   string `json:"type"`
				} `json:"embedHeight"`
				EmbedHtml struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"embedHtml"`
				EmbedWidth struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"embedWidth"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoPlayer"`
		VideoProcessingDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				EditorSuggestionsAvailability struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"editorSuggestionsAvailability"`
				FileDetailsAvailability struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"fileDetailsAvailability"`
				ProcessingFailureReason struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"processingFailureReason"`
				ProcessingIssuesAvailability struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"processingIssuesAvailability"`
				ProcessingProgress struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"processingProgress"`
				ProcessingStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"processingStatus"`
				TagSuggestionsAvailability struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"tagSuggestionsAvailability"`
				ThumbnailsAvailability struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"thumbnailsAvailability"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoProcessingDetails"`
		VideoProcessingDetailsProcessingProgress struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				PartsProcessed struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"partsProcessed"`
				PartsTotal struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"partsTotal"`
				TimeLeftMs struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"timeLeftMs"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoProcessingDetailsProcessingProgress"`
		VideoProjectDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoProjectDetails"`
		VideoRating struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Rating struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"rating"`
				VideoId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"videoId"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoRating"`
		VideoRecordingDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Location struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"location"`
				LocationDescription struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"locationDescription"`
				RecordingDate struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"recordingDate"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoRecordingDetails"`
		VideoSnippet struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CategoryId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"categoryId"`
				ChannelId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelId"`
				ChannelTitle struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"channelTitle"`
				DefaultAudioLanguage struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"defaultAudioLanguage"`
				DefaultLanguage struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"defaultLanguage"`
				Description struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"description"`
				LiveBroadcastContent struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"liveBroadcastContent"`
				Localized struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"localized"`
				PublishedAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishedAt"`
				Tags struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"tags"`
				Thumbnails struct {
					Ref         string `json:"$ref"`
					Description string `json:"description"`
				} `json:"thumbnails"`
				Title struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"title"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoSnippet"`
		VideoStatistics struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CommentCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"commentCount"`
				DislikeCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"dislikeCount"`
				FavoriteCount struct {
					Deprecated  bool   `json:"deprecated"`
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"favoriteCount"`
				LikeCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"likeCount"`
				ViewCount struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"viewCount"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoStatistics"`
		VideoStatus struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				Embeddable struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"embeddable"`
				FailureReason struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"failureReason"`
				License struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"license"`
				MadeForKids struct {
					Type string `json:"type"`
				} `json:"madeForKids"`
				PrivacyStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"privacyStatus"`
				PublicStatsViewable struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"publicStatsViewable"`
				PublishAt struct {
					Description string `json:"description"`
					Format      string `json:"format"`
					Type        string `json:"type"`
				} `json:"publishAt"`
				RejectionReason struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"rejectionReason"`
				SelfDeclaredMadeForKids struct {
					Type string `json:"type"`
				} `json:"selfDeclaredMadeForKids"`
				UploadStatus struct {
					Description      string   `json:"description"`
					Enum             []string `json:"enum"`
					EnumDescriptions []string `json:"enumDescriptions"`
					Type             string   `json:"type"`
				} `json:"uploadStatus"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoStatus"`
		VideoSuggestions struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				EditorSuggestions struct {
					Description string `json:"description"`
					Items       struct {
						Enum             []string `json:"enum"`
						EnumDescriptions []string `json:"enumDescriptions"`
						Type             string   `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"editorSuggestions"`
				ProcessingErrors struct {
					Description string `json:"description"`
					Items       struct {
						Enum             []string `json:"enum"`
						EnumDescriptions []string `json:"enumDescriptions"`
						Type             string   `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"processingErrors"`
				ProcessingHints struct {
					Description string `json:"description"`
					Items       struct {
						Enum             []string `json:"enum"`
						EnumDescriptions []string `json:"enumDescriptions"`
						Type             string   `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"processingHints"`
				ProcessingWarnings struct {
					Description string `json:"description"`
					Items       struct {
						Enum             []string `json:"enum"`
						EnumDescriptions []string `json:"enumDescriptions"`
						Type             string   `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"processingWarnings"`
				TagSuggestions struct {
					Description string `json:"description"`
					Items       struct {
						Ref string `json:"$ref"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"tagSuggestions"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoSuggestions"`
		VideoSuggestionsTagSuggestion struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				CategoryRestricts struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"categoryRestricts"`
				Tag struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"tag"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoSuggestionsTagSuggestion"`
		VideoTopicDetails struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				RelevantTopicIds struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"relevantTopicIds"`
				TopicCategories struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"topicCategories"`
				TopicIds struct {
					Description string `json:"description"`
					Items       struct {
						Type string `json:"type"`
					} `json:"items"`
					Type string `json:"type"`
				} `json:"topicIds"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"VideoTopicDetails"`
		WatchSettings struct {
			Description string `json:"description"`
			Id          string `json:"id"`
			Properties  struct {
				BackgroundColor struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"backgroundColor"`
				FeaturedPlaylistId struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"featuredPlaylistId"`
				TextColor struct {
					Description string `json:"description"`
					Type        string `json:"type"`
				} `json:"textColor"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"WatchSettings"`
	} `json:"schemas"`
	ServicePath string `json:"servicePath"`
	Title       string `json:"title"`
	Version     string `json:"version"`
}
