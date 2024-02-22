-- Copyright 2024 Specter Ops, Inc.
--
-- Licensed under the Apache License, Version 2.0
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     http://www.apache.org/licenses/LICENSE-2.0
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.
--
-- SPDX-License-Identifier: Apache-2.0

INSERT INTO permissions (authority, name, created_at, updated_at) VALUES ('db', 'Wipe', NOW(), NOW());

INSERT INTO roles_permissions (role_id, permission_id) 
VALUES ((SELECT id FROM roles where name = 'Administrator'), 
(SELECT id FROM permissions where authority='db' AND name = 'Wipe'));