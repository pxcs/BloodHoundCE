// Copyright 2024 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { useCallback, useEffect, useState } from 'react';
import { getSelfResponse } from 'src/ducks/auth/types';
import { PermissionsSpec } from 'bh-shared-ui';
import { useAppSelector } from 'src/store';

type PermissionState = {
    hasAtLeastOne: boolean;
    hasAll: boolean;
};

const usePermissions: (permissions: PermissionsSpec[]) => PermissionState = (permissions) => {
    const [permissionState, setPermissionState] = useState<PermissionState>({ hasAtLeastOne: false, hasAll: false });
    const auth = useAppSelector((state) => state.auth);

    const checkUserPermissions = useCallback(
        (user: getSelfResponse): PermissionState => {
            const userPermissions = user.roles.map((role) => role.permissions).flat();
            let hasAtLeastOne = false;

            const hasAll = permissions.every((permission) => {
                return userPermissions.some((userPermission) => {
                    const matched =
                        userPermission.authority === permission.authority && userPermission.name === permission.name;

                    if (matched && !hasAtLeastOne) hasAtLeastOne = true;
                    return matched;
                });
            });

            return { hasAtLeastOne, hasAll };
        },
        [permissions]
    );

    useEffect(() => {
        if (auth.user) {
            setPermissionState(checkUserPermissions(auth.user));
        }
    }, [auth, checkUserPermissions]);

    return permissionState;
};

export default usePermissions;
