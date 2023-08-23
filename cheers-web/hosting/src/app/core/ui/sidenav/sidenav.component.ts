import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {firstValueFrom} from "rxjs";

@Component({
    selector: 'app-sidenav',
    templateUrl: './sidenav.component.html',
    styleUrls: ['./sidenav.component.sass']
})
export class SidenavComponent implements OnInit {

    isBusiness: boolean = false

    businessItems = [
        {
            title: "Payouts",
            icon: "account_balance",
            routerLink: "/business/payouts",
        },
        {
            title: "Orders",
            icon: "receipt_long",
            routerLink: "/business/orders",
        },
        {
            title: "Reports",
            icon: "candlestick_chart",
            routerLink: "/business/reports",
        },
    ]

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
            icon: "celebration",
            routerLink: "/parties",
        },
        {
            title: "Tickets",
            icon: "local_activity",
            routerLink: "/tickets",
        },
        {
            title: "Map",
            icon: "map",
            routerLink: "/map",
        },
        {
            title: "Chats",
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
        public auth: AngularFireAuth,
    ) {
    }

    async ngOnInit()  {
        const token = await firstValueFrom(this.auth.idTokenResult)
        if (token == null)
            return
        this.isBusiness = token.claims["business"] != null
    }

}
