import logo from '$lib/images/logo.svg';
import pictureDefault from '$lib/images/profile.png'

export const linkUtil = {
    logo: true,
    logoSrc: logo,
    logoLink: true,
    linkUrl: '/',
    altText: 'Logo',
    links: [
        {
            id: 'profile',
            url: '/profile',
            displayInNav: true,
            displayInFooter: true,
            linkText: "",
            picture: pictureDefault
        },
    ]
}