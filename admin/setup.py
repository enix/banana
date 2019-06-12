#! /usr/bin/env python3

import setuptools

setuptools.setup(
    name='bananadm',
    version='0.1.5',
    author='Arthur Chaloin from Enix <enix.io>',
    author_email='arthur.chaloin@enix.fr',
    description='A command line tool to manage Banana instances',
    packages=setuptools.find_packages(),
    scripts=['bananadm'],
    classifiers=[
        'Programming Language :: Python :: 3',
        'Operating System :: OS Independent',
    ],
    install_requires=[
        'urllib3',
        'requests>=2.21.0',
        'hvac',
    ],
)
