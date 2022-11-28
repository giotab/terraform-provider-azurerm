package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageProfile struct {
	DataDisks      *[]DataDisk     `json:"dataDisks,omitempty"`
	ImageReference *ImageReference `json:"imageReference"`
	OsDisk         *OSDisk         `json:"osDisk"`
}
