/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Joseph Martinsen <joseph@martinsen.com>
 */

import { useState } from "react";
import "./Reciept.css"

function Reciept() {
    const [title, updateTitle] = useState("Location");
    return <div className="reciept">
        <h1 className="title">
            GoSplitsies:<input className="reciept-title" value={title} onChange={e => updateTitle(e.target.value)} />
        </h1>
    </div>
}

export default Reciept;
