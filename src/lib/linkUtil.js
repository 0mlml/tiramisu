import logo from './images/logo.svg';
import profilePhoto from './images/profile.png'

const name ='a'

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