/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Joseph Martinsen <joseph@martinsen.com>
 */

import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
