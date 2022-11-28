package disks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskUpdateProperties struct {
	BurstingEnabled              *bool                         `json:"burstingEnabled,omitempty"`
	DataAccessAuthMode           *DataAccessAuthMode           `json:"dataAccessAuthMode,omitempty"`
	DiskAccessId                 *string                       `json:"diskAccessId,omitempty"`
	DiskIOPSReadOnly             *int64                        `json:"diskIOPSReadOnly,omitempty"`
	DiskIOPSReadWrite            *int64                        `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadOnly             *int64                        `json:"diskMBpsReadOnly,omitempty"`
	DiskMBpsReadWrite            *int64                        `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeGB                   *int64                        `json:"diskSizeGB,omitempty"`
	Encryption                   *Encryption                   `json:"encryption"`
	EncryptionSettingsCollection *EncryptionSettingsCollection `json:"encryptionSettingsCollection"`
	MaxShares                    *int64                        `json:"maxShares,omitempty"`
	NetworkAccessPolicy          *NetworkAccessPolicy          `json:"networkAccessPolicy,omitempty"`
	OsType                       *OperatingSystemTypes         `json:"osType,omitempty"`
	PropertyUpdatesInProgress    *PropertyUpdatesInProgress    `json:"propertyUpdatesInProgress"`
	PublicNetworkAccess          *PublicNetworkAccess          `json:"publicNetworkAccess,omitempty"`
	PurchasePlan                 *PurchasePlan                 `json:"purchasePlan"`
	SupportedCapabilities        *SupportedCapabilities        `json:"supportedCapabilities"`
	SupportsHibernation          *bool                         `json:"supportsHibernation,omitempty"`
	Tier                         *string                       `json:"tier,omitempty"`
}
