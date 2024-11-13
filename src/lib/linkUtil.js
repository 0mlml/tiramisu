import logo from './images/logo.svg';
import profilePhoto from './images/profile.png'
import * as fetch from "./fetcher"

const token = (await fetch.loginUser("user@example.com", "user123!@#"))['data']['token'];

const profile = await fetch.getProfile(token);

const name = profile['data']['name'];

export const linkUtil = {
    logo: true,
    logoSrc: logo,
    logoLink: true,
    linkUrl: ' ',
    altText: 'Logo',
    links:[
    {id:'profile',
    url:'#/profile',
    displayInNav: true,
    displayInFooter: true,
    linkText: name,
    picture: profilePhoto
    },
    ]
}