/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2016 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package client

import (
	"context"
	"net/http"
	"net/url"

	v1 "go.k6.io/k6/api/v1"
)

// Status returns the current k6 status.
func (c *Client) Status(ctx context.Context) (ret v1.Status, err error) {
	var resp v1.StatusJSONAPI

	if err = c.CallAPI(ctx, http.MethodGet, &url.URL{Path: "/v1/status"}, nil, &resp); err != nil {
		return ret, err
	}

	return resp.Status(), nil
}

// SetStatus tries to change the current status and returns the new one if it
// was successful.
func (c *Client) SetStatus(ctx context.Context, patch v1.Status) (ret v1.Status, err error) {
	var resp v1.StatusJSONAPI

	apiURL := &url.URL{Path: "/v1/status"}
	if err = c.CallAPI(ctx, http.MethodPatch, apiURL, v1.NewStatusJSONAPI(patch), &resp); err != nil {
		return ret, err
	}

	return resp.Status(), nil
}
