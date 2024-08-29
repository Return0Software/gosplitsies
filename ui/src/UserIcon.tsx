/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Joseph Martinsen <joseph@martinsen.com>
 */

import "./UserIcon.css"

function DefaultIcon() {
    return (
        <div className="default-avatar">
            {/* TODO: {profile.name.charAt(0).toUpperCase()} */}
            J
        </div>
    )
}

function ImgUserIcon() {
    return <img className="user-icon-outline" src="https://picsum.photos/50" />
}


export type UserIconProps = {
    type: "img"
    img?: string;
    onRemove: (item: UserIconProps) => void;
} | { type: "default" | "add"}

function UserIcon(props: UserIconProps) {
    switch(props.type) {
        case "default":
            return <DefaultIcon />;
        case "img":
            return <ImgUserIcon />
        case "add":
            return <div>
                <p>ADD</p>
            </div>
    }
}

export default UserIcon;