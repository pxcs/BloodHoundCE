// Copyright 2023 Specter Ops, Inc.
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

// Definitions for third-party modules with no associated types
declare module 'dagrejs';

declare module '@neo4j-cypher/react-codemirror' {
    type CypherEditorProps = import('@neo4j-cypher/react-codemirror/src/react-codemirror.d.ts').CypherEditorProps;
    export class CypherEditor extends React.Component<CypherEditorProps, any> {}
}
