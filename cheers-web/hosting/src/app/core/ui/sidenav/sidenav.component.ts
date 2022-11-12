import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";

@Component({
    selector: 'app-sidenav',
    templateUrl: './sidenav.component.html',
    styleUrls: ['./sidenav.component.sass']
})
export class SidenavComponent implements OnInit {

    more_items = [
        {
            title: "Settings",
            icon: "settings",
            routerLink: "/settings",
        },
        {
            title: "Help",
            icon: "help",
            routerLink: "/help",
        },
        {
            title: "Send feedback",
            icon: "info",
            routerLink: "/send-feedback",
        }
    ]

    items = [
        {
            title: "Home",
            icon: "home",
            routerLink: "/home",
        },
        {
            title: "Parties",
            icon: "explore",
            routerLink: "/parties/feed",
        },
        {
            title: "Tickets",
            icon: "local_activity",
            routerLink: "/tickets",
        },
        {
            title: "Messages",
            icon: "message",
            routerLink: "/chats",
        },
        {
            title: "Notifications",
            icon: "notifications",
            routerLink: "/notifications",
        },
        {
            title: "Profile",
            icon: "account_circle",
            routerLink: "/profile",
        },
    ]

    constructor(
    ) {
    }

    ngOnInit(): void {
    }

}
