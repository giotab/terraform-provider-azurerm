package videoanalyzers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VideoAnalyzerPropertiesUpdate struct {
	Encryption      *AccountEncryption `json:"encryption"`
	Endpoints       *[]Endpoint        `json:"endpoints,omitempty"`
	StorageAccounts *[]StorageAccount  `json:"storageAccounts,omitempty"`
}
