/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Joseph Martinsen <joseph@martinsen.com>
 */

import { useReducer } from "react";
import "./FacePile.css";
import UserIcon, { UserIconProps } from "./UserIcon";

type SplitUser  = Omit<UserIconProps, "onRemove">;

function splitUserReducer(state: SplitUser[], action: {type: "add" | "remove", item: SplitUser}) {
    switch (action.type) {
        case "add":
            return [...state, action.item];
        case "remove":
            return state.filter(item => item !== action.item)
        default:
            return state;
    }
}

function useSplitUser() {
    return useReducer(splitUserReducer, []);
}

function FacePile() {
    const [splitUsers, dispatchSplitUser] = useSplitUser();

    function removeSplitUser(user: SplitUser) {
        dispatchSplitUser({ type: "remove", item: user})
    }

    return (
        <div className="header sticky">
            <div className="profile-list-container">
                <div className="profile-list">
                    {splitUsers.map((splitUser, index) => (
                        <div className="profile-item" key={index}>
                            <UserIcon {...splitUser} onRemove={removeSplitUser} />
                        </div>
                    ))}
                </div>
            </div>
            <button onClick={() => {
                dispatchSplitUser({ type: "add", item: splitUsers.length % 2 ? { type: "img" } : { type: "default" }})
            }}>Add</button>
        </div>
    )
}

export default FacePile;