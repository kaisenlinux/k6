// Package extensions provides the Chrome DevTools Protocol
// commands, types, and events for the Extensions domain.
//
// Defines commands and events for browser extensions.
//
// Generated by the cdproto-gen command.
package extensions

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/mailru/easyjson"
)

// LoadUnpackedParams installs an unpacked extension from the filesystem
// similar to --load-extension CLI flags. Returns extension ID once the
// extension has been installed. Available if the client is connected using the
// --remote-debugging-pipe flag and the --enable-unsafe-extension-debugging flag
// is set.
type LoadUnpackedParams struct {
	Path string `json:"path"` // Absolute file path.
}

// LoadUnpacked installs an unpacked extension from the filesystem similar to
// --load-extension CLI flags. Returns extension ID once the extension has been
// installed. Available if the client is connected using the
// --remote-debugging-pipe flag and the --enable-unsafe-extension-debugging flag
// is set.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Extensions#method-loadUnpacked
//
// parameters:
//
//	path - Absolute file path.
func LoadUnpacked(path string) *LoadUnpackedParams {
	return &LoadUnpackedParams{
		Path: path,
	}
}

// LoadUnpackedReturns return values.
type LoadUnpackedReturns struct {
	ID string `json:"id,omitempty"` // Extension id.
}

// Do executes Extensions.loadUnpacked against the provided context.
//
// returns:
//
//	id - Extension id.
func (p *LoadUnpackedParams) Do(ctx context.Context) (id string, err error) {
	// execute
	var res LoadUnpackedReturns
	err = cdp.Execute(ctx, CommandLoadUnpacked, p, &res)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}

// GetStorageItemsParams gets data from extension storage in the given
// storageArea. If keys is specified, these are used to filter the result.
type GetStorageItemsParams struct {
	ID          string      `json:"id"`             // ID of extension.
	StorageArea StorageArea `json:"storageArea"`    // StorageArea to retrieve data from.
	Keys        []string    `json:"keys,omitempty"` // Keys to retrieve.
}

// GetStorageItems gets data from extension storage in the given storageArea.
// If keys is specified, these are used to filter the result.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Extensions#method-getStorageItems
//
// parameters:
//
//	id - ID of extension.
//	storageArea - StorageArea to retrieve data from.
func GetStorageItems(id string, storageArea StorageArea) *GetStorageItemsParams {
	return &GetStorageItemsParams{
		ID:          id,
		StorageArea: storageArea,
	}
}

// WithKeys keys to retrieve.
func (p GetStorageItemsParams) WithKeys(keys []string) *GetStorageItemsParams {
	p.Keys = keys
	return &p
}

// GetStorageItemsReturns return values.
type GetStorageItemsReturns struct {
	Data easyjson.RawMessage `json:"data,omitempty"`
}

// Do executes Extensions.getStorageItems against the provided context.
//
// returns:
//
//	data
func (p *GetStorageItemsParams) Do(ctx context.Context) (data easyjson.RawMessage, err error) {
	// execute
	var res GetStorageItemsReturns
	err = cdp.Execute(ctx, CommandGetStorageItems, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

// RemoveStorageItemsParams removes keys from extension storage in the given
// storageArea.
type RemoveStorageItemsParams struct {
	ID          string      `json:"id"`          // ID of extension.
	StorageArea StorageArea `json:"storageArea"` // StorageArea to remove data from.
	Keys        []string    `json:"keys"`        // Keys to remove.
}

// RemoveStorageItems removes keys from extension storage in the given
// storageArea.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Extensions#method-removeStorageItems
//
// parameters:
//
//	id - ID of extension.
//	storageArea - StorageArea to remove data from.
//	keys - Keys to remove.
func RemoveStorageItems(id string, storageArea StorageArea, keys []string) *RemoveStorageItemsParams {
	return &RemoveStorageItemsParams{
		ID:          id,
		StorageArea: storageArea,
		Keys:        keys,
	}
}

// Do executes Extensions.removeStorageItems against the provided context.
func (p *RemoveStorageItemsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandRemoveStorageItems, p, nil)
}

// ClearStorageItemsParams clears extension storage in the given storageArea.
type ClearStorageItemsParams struct {
	ID          string      `json:"id"`          // ID of extension.
	StorageArea StorageArea `json:"storageArea"` // StorageArea to remove data from.
}

// ClearStorageItems clears extension storage in the given storageArea.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Extensions#method-clearStorageItems
//
// parameters:
//
//	id - ID of extension.
//	storageArea - StorageArea to remove data from.
func ClearStorageItems(id string, storageArea StorageArea) *ClearStorageItemsParams {
	return &ClearStorageItemsParams{
		ID:          id,
		StorageArea: storageArea,
	}
}

// Do executes Extensions.clearStorageItems against the provided context.
func (p *ClearStorageItemsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandClearStorageItems, p, nil)
}

// SetStorageItemsParams sets values in extension storage in the given
// storageArea. The provided values will be merged with existing values in the
// storage area.
type SetStorageItemsParams struct {
	ID          string              `json:"id"`          // ID of extension.
	StorageArea StorageArea         `json:"storageArea"` // StorageArea to set data in.
	Values      easyjson.RawMessage `json:"values"`
}

// SetStorageItems sets values in extension storage in the given storageArea.
// The provided values will be merged with existing values in the storage area.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Extensions#method-setStorageItems
//
// parameters:
//
//	id - ID of extension.
//	storageArea - StorageArea to set data in.
//	values - Values to set.
func SetStorageItems(id string, storageArea StorageArea, values easyjson.RawMessage) *SetStorageItemsParams {
	return &SetStorageItemsParams{
		ID:          id,
		StorageArea: storageArea,
		Values:      values,
	}
}

// Do executes Extensions.setStorageItems against the provided context.
func (p *SetStorageItemsParams) Do(ctx context.Context) (err error) {
	return cdp.Execute(ctx, CommandSetStorageItems, p, nil)
}

// Command names.
const (
	CommandLoadUnpacked       = "Extensions.loadUnpacked"
	CommandGetStorageItems    = "Extensions.getStorageItems"
	CommandRemoveStorageItems = "Extensions.removeStorageItems"
	CommandClearStorageItems  = "Extensions.clearStorageItems"
	CommandSetStorageItems    = "Extensions.setStorageItems"
)